package dnsProviders

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

// function that creates a new route53 session
func NewRoute53Session() (*session.Session, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// function that creates a client for Route53
func NewRoute53Client() (*route53.Route53, error) {
	s, err := NewRoute53Session()
	if err != nil {
		return nil, err
	}
	return route53.New(s), nil
}

// function that gets the ZoneId based on the DNS Name
func GetZoneId(c *route53.Route53, hostname string) (string, error) {
	// extract root domain name from full hostname using spliting string after second dot
	rootDomain := strings.Split(hostname, ".")[1] + "." + strings.Split(hostname, ".")[2]

	// retrieve zone id from route 53 using session and hostname
	params := &route53.ListHostedZonesByNameInput{
		DNSName: aws.String(rootDomain),
	}
	hostedZone, err := c.ListHostedZonesByName(params)
	if err != nil {
		return "", err
	}
	// check the length of the hostedzone.hostedzones array
	if len(hostedZone.HostedZones) == 0 {
		return "", fmt.Errorf("no hosted zone found for %s", rootDomain)
	}
	return *hostedZone.HostedZones[0].Id, nil
}

// function that updates A or AAAA record in Route53 zone
func UpdateRecord(svc *route53.Route53, hostname string, target string, recordType string, TTL int64, zoneId string) error {
	// fmt.Println("Updating A record with following parameters:")
	// fmt.Printf("Hostname: %s\nTarget: %s\nRecordType: %s\nTTL: %d\nZoneId: %s\n", hostname, target, recordType, TTL, zoneId)
	comment := "Change performed at " + time.Now().String()

	// if recordType is not A or AAAA then throw an error
	if recordType != "A" && recordType != "AAAA" {
		return fmt.Errorf("record type: %s not supported", recordType)
	}
	// update the dns record
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(hostname),   // Required
						Type: aws.String(recordType), // Required A or AAAA
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(target), // Required - IP address
							},
						},
						TTL: aws.Int64(TTL | 600),
						// SetIdentifier: aws.String("Simple"), // Optional when it is other than simple
					},
				},
			},
			Comment: aws.String(comment),
		},
		HostedZoneId: aws.String(zoneId), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		return err
	}

	// Pretty-print the response data.
	fmt.Println("Successfully updated DNS record")
	fmt.Println(resp)
	return nil
}
