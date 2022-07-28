package updater

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"

	dnsProviders "go.pkg.dipak.io/ddns-server/internal/dns-providers"
)

type Updater struct{}

var domainRegexp = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`)

func CheckIPAddressType(ip string) (string, error) {
	for i := 0; i < len(ip); i++ {
		switch ip[i] {
		case '.':
			return "ipv4", nil
		case ':':
			return "ipv6", nil
		}
	}
	return "", fmt.Errorf("IP address is not valid")
}

// function that validates the ip
func ValidateIP(ip string) bool {
	// validate the ip
	if net.ParseIP(ip) == nil {
		return false // ip is not valid or unable to parse
	} else {
		return true // ip is valid
	}
}

func IsValidDomain(domain string) bool {
	return domainRegexp.MatchString(domain)
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

// function that validates the hostname
func ValidateHostname(hostname string) bool {
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

func UpdateDNSRecord(hostname string, ip string, TTL int64) error {
	ip_type, err := CheckIPAddressType(ip)
	if err != nil {
		return err
	}

	c, err := dnsProviders.NewRoute53Client()
	if err != nil {
		fmt.Println("failed to create client,", err)
		return err
	}

	zoneId, err := dnsProviders.GetZoneId(c, hostname)

	if ip_type == "ipv4" {
		dnsProviders.UpdateRecord(c, hostname, ip, "A", TTL, zoneId)
	} else if ip_type == "ipv6" {
		dnsProviders.UpdateRecord(c, hostname, ip, "AAAA", TTL, zoneId)
	} else {
		log.Println("IP address is not valid")
	}
	return err

}

// function that handles dns update requests with params in the url hostname and myip
func UpdateDNS(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	// if the username and password are not valid, return false
	if !ok || username != "username" || password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	// get the hostname and ip from the url
	hostname := r.URL.Query().Get("hostname")
	ip := r.URL.Query().Get("myip")
	validRequest := ValidateHostname(hostname) && ValidateIP(ip)
	if validRequest {
		log.Println("hostname:", hostname, "ip:", ip)

		// update the dns record
		err := UpdateDNSRecord(hostname, ip, 600)
		if err != nil {
			log.Println("Error updating dns record:", err)
		}
		// return a success message
		w.Write([]byte("OK"))
	} else {
		// return a failure message
		w.Write([]byte("Failed to validate the request. Please check following:"))
		if !ValidateHostname(hostname) {
			w.Write([]byte("\nhostname is not valid"))
		}
		if !ValidateIP(ip) {
			w.Write([]byte("\nip is not valid"))
		}
	}
}
