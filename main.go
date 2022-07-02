package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var domainRegexp = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`)

var name string
var target string
var TTL int64
var weight = int64(1)
var zoneId string = "" // update this manually

func main() {
	log.Println("Starting server at 127.0.0.1:80")
	// err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/ddns.mydomain.net/fullchain.pem", "/etc/letsencrypt/live/ddns.mydomain.net/privkey.pem", &handler{})
	// err := http.ListenAndServe(":80", &handler{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/nic/update", updateDNS)

	log.Fatal(http.ListenAndServe(":80", nil))

}

// type rootHandler struct{}

func rootHandler(resp http.ResponseWriter, req *http.Request) {

	// log.Println(req.URL.String())
	auth, found := req.Header["Authorization"]
	if found {
		log.Println("Authorization:", auth)
	}
	agent, found := req.Header["User-Agent"]
	if found {
		log.Println("User-Agent:", agent)
	}

	if req.URL.Path == "/" {
		h := resp.Header()
		h.Set("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Server is running okay."))
	} else {
		// return a 404
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("Not found."))
	}

}

// function that handles dns update requests with params in the url hostname and myip
func updateDNS(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	// if the username and password are not valid, return false
	if !ok || username != "username" || password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// get the hostname and ip from the url
	hostname := r.URL.Query().Get("hostname")
	ip := r.URL.Query().Get("myip")
	validRequest := validateHostname(hostname) && validateIP(ip)
	if validRequest {
		log.Println("hostname:", hostname, "ip:", ip)
		name = hostname
		target = ip
		TTL = 3600
		// update the dns record
		err := updateDNSRecord(hostname, ip)
		if err != nil {
			log.Println("Error updating dns record:", err)
		}
		// return a success message
		w.Write([]byte("OK"))
	} else {
		// return a failure message
		w.Write([]byte("Falied to validate the request. Please check follwing:"))
		if !validateHostname(hostname) {
			w.Write([]byte("\nhostname is not valid"))
		}
		if !validateIP(ip) {
			w.Write([]byte("\nip is not valid"))
		}
	}
}

func updateDNSRecord(hostname string, ip string) error {
	// update the dns record
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return err
	}

	svc := route53.New(sess)
	// log.Println(sess.Config.Credentials.Get())

	if checkIPAddressType(ip) == "ipv4" {
		updateARecord(svc)
	} else if checkIPAddressType(ip) == "ipv6" {
		updateAAAARecord(svc)
	} else {
		log.Println("IP address is not valid")
	}
	return err

}

// NormalizeDomain returns a normalized domain.
// It returns an empty string if the domain is not valid.
func NormalizeDomain(domain string) string {
	// Trim whitespace.
	domain = strings.TrimSpace(domain)
	// Check validity.
	if !IsValidDomain(domain) {
		return ""
	}
	// Remove trailing dot.
	domain = strings.TrimRight(domain, ".")
	// Convert to lower case.
	domain = strings.ToLower(domain)
	return domain
}

func IsValidDomain(domain string) bool {
	return domainRegexp.MatchString(domain)
}

// function that validates the hostname
func validateHostname(hostname string) bool {
	// validate the hostname
	if hostname == "" {
		return false // empty hostname
	} else {
		if NormalizeDomain(hostname) != "" {
			return true // valid hostname
		} else {
			return false // hostname is not valid
		}
	}
}

// function that validates the ip
func validateIP(ip string) bool {
	// validate the ip
	if net.ParseIP(ip) == nil {
		return false // ip is not valid or unable to parse
	} else {
		return true // ip is valid
	}
}

func checkIPAddressType(ip string) string {
	for i := 0; i < len(ip); i++ {
		switch ip[i] {
		case '.':
			return "ipv4"
		case ':':
			return "ipv6"
		}
	}
	return ""
}

// function that updates A or AAAA record in Route53 zone
func updateARecord(svc *route53.Route53) error {
	// update the dns record
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(name), // Required
						Type: aws.String("A"),  // Required
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(target), // Required - IP address
							},
						},
						TTL:           aws.Int64(TTL),
						SetIdentifier: aws.String("Arbitrary Id describing this change set"),
					},
				},
			},
			Comment: aws.String("Sample update."),
		},
		HostedZoneId: aws.String(zoneId), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	fmt.Println("Change Response:")
	fmt.Println(resp)
	return nil
}

func updateAAAARecord(svc *route53.Route53) error {
	// update the dns record
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(name),   // Required
						Type: aws.String("AAAA"), // Required
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(target), // Required - IP address
							},
						},
						TTL:           aws.Int64(TTL),
						SetIdentifier: aws.String("Arbitrary Id describing this change set"),
					},
				},
			},
			Comment: aws.String("Sample update."),
		},
		HostedZoneId: aws.String(zoneId), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	fmt.Println("Change Response:")
	fmt.Println(resp)
	return nil
}

// function that gets the ZoneId based on the username and hostname
func getZoneId(svc *route53.Route53, username string, hostname string) (string, error) {
	var zoneId string
	// get the zone id (need to use database or make a request to get the zone id)
	if username == "username" && hostname == "hostname" {
		zoneId = "Z2M3LMPEXAMPLE"
	}

	return zoneId, nil
}
