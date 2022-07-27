package dnsProviders

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func TestNewRoute53Session(t *testing.T) {
	tests := []struct {
		name    string
		want    *session.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoute53Session()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoute53Session() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoute53Session() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRoute53Client(t *testing.T) {
	tests := []struct {
		name    string
		want    *route53.Route53
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoute53Client()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoute53Client() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoute53Client() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetZoneId(t *testing.T) {
	type args struct {
		c        *route53.Route53
		hostname string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetZoneId(tt.args.c, tt.args.hostname)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetZoneId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetZoneId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	type args struct {
		svc        *route53.Route53
		hostname   string
		target     string
		recordType string
		TTL        int64
		zoneId     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateRecord(tt.args.svc, tt.args.hostname, tt.args.target, tt.args.recordType, tt.args.TTL, tt.args.zoneId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
