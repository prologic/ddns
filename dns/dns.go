package dns

import (
	"errors"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/miekg/dns"
	"github.com/muka/dyndns/db"
)

//GetKey return the reverse domain
func GetKey(domain string, rtype uint16) (r string, e error) {
	log.Debugf("Get key for %s", domain)
	if n, ok := dns.IsDomainName(domain); ok {

		labels := dns.SplitDomainName(domain)

		// Reverse domain, starting from top-level domain
		// eg.  ".com.mkaczanowski.test "
		var tmp string
		for i := 0; i < int(math.Floor(float64(n/2))); i++ {
			tmp = labels[i]
			labels[i] = labels[n-1]
			labels[n-1] = tmp
		}

		reverseDomain := strings.Join(labels, ".")
		r = strings.Join([]string{reverseDomain, strconv.Itoa(int(rtype))}, "_")
		log.Debugf("Key is %s", r)
	} else {
		e = errors.New("Invalid domain:  " + domain)
		log.Error(e.Error())
	}

	return r, e
}

//GetRecord return a new DNS record
func GetRecord(domain string, rtype uint16) (dns.RR, error) {

	log.Debugf("Load record %s", domain)

	key, err := GetKey(domain, rtype)
	if err != nil {
		return nil, err
	}

	record, err := db.GetRecord(key)
	if err == nil {
		return dns.NewRR(record.RR)
	}

	return nil, err
}

//UpdateRecord update or remove a record
func UpdateRecord(r dns.RR, q *dns.Question) error {

	var (
		rr    dns.RR
		name  string
		rtype uint16
		ttl   uint32
		ip    net.IP
	)

	header := r.Header()
	name = header.Name
	rtype = header.Rrtype
	ttl = header.Ttl

	revName, err := GetKey(name, rtype)

	if err != nil {
		return err
	}

	if _, ok := dns.IsDomainName(name); ok {
		if header.Class == dns.ClassANY && header.Rdlength == 0 { // Delete record
			log.Debugf("Remove %s", name)
			db.DeleteRecord(revName)
		} else {

			// Add record
			rheader := GetHeader(name, rtype, ttl)

			if a, ok := r.(*dns.A); ok {
				rrr, err := GetRecord(name, rtype)
				if err == nil {
					rr = rrr.(*dns.A)
				} else {
					rr = new(dns.A)
				}

				ip = a.A
				rr.(*dns.A).Hdr = rheader
				rr.(*dns.A).A = ip

			} else if a, ok := r.(*dns.AAAA); ok {

				rrr, err := GetRecord(name, rtype)
				if err == nil {
					rr = rrr.(*dns.AAAA)
				} else {
					rr = new(dns.AAAA)
				}

				ip = a.AAAA
				rr.(*dns.AAAA).Hdr = rheader
				rr.(*dns.AAAA).AAAA = ip

			}

			rrKey, err1 := GetKey(rr.Header().Name, rr.Header().Rrtype)
			if err1 != nil {
				return err1
			}

			log.Debugf("Saving record %s (%s)", rr.Header().Name, rrKey)

			record := db.NewRecord(rr.String(), 0)
			db.StoreRecord(rrKey, record)

		}
	}

	return nil
}

// GetHeader create a new record header
func GetHeader(name string, rtype uint16, ttl uint32) dns.RR_Header {
	return dns.RR_Header{
		Name:   name,
		Rrtype: rtype,
		Class:  dns.ClassINET,
		Ttl:    ttl,
	}
}

//AddPTRRecord for the specified domain and ip address
func AddPTRRecord(ip string, domain string, ttl uint32, expires int64) error {

	rtype := dns.TypePTR

	rr := new(dns.PTR)
	rr.Ptr = ip
	rr.Hdr = GetHeader(domain, rtype, ttl)

	key, err := GetKey(domain, rtype)
	if err != nil {
		return err
	}

	log.Debugf("Adding PTR Record %s > %s", ip, domain)
	record := db.NewRecord(rr.String(), expires)

	return db.StoreRecord(key, record)
}

func parseQuery(m *dns.Msg) {
	var rr dns.RR
	for _, q := range m.Question {
		log.Debugf("DNS query: %s", q.String())
		readRR, e := GetRecord(q.Name, q.Qtype)
		if e != nil {
			log.Errorf("Error getting record: %s", e.Error())
			continue
		}
		rr = readRR.(dns.RR)
		if rr.Header().Name == q.Name {
			log.Debugf("Found match: %s", rr.String())
			m.Answer = append(m.Answer, rr)
		}

	}
}

//HandleDNSRequest handle incoming requests
func HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg, enableUpdates bool) {

	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		log.Debugf("Got query request")
		parseQuery(m)

	case dns.OpcodeUpdate:
		if enableUpdates {
			log.Debugf("Got update request")
			for _, question := range r.Question {
				for _, rr := range r.Ns {
					UpdateRecord(rr, &question)
				}
			}
		} else {
			log.Debugf("Update request ignored, TSIG missing")
		}
	}

	if r.IsTsig() != nil {
		if w.TsigStatus() == nil {
			m.SetTsig(r.Extra[len(r.Extra)-1].(*dns.TSIG).Hdr.Name,
				dns.HmacMD5, 300, time.Now().Unix())
		} else {
			log.Println("Status ", w.TsigStatus().Error())
		}
	}

	w.WriteMsg(m)
}

//Serve the DNS server
func Serve(name, secret string, port int) error {

	log.Debugf("Starting server on :%d", port)
	server := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}

	if name != "" {
		server.TsigSecret = map[string]string{name: secret}
	}

	err := server.ListenAndServe()
	defer server.Shutdown()

	if err != nil {
		log.Fatalf("Failed to setup the udp server: %s", err.Error())
	}

	return err
}
