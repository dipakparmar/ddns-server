package updater

import (
	"net/http"
	"testing"
)

func TestCheckIPAddressType(t *testing.T) {
	type args struct {
		ip string
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
			got, err := CheckIPAddressType(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIPAddressType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIPAddressType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIP(tt.args.ip); got != tt.want {
				t.Errorf("ValidateIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDomain(tt.args.domain); got != tt.want {
				t.Errorf("IsValidDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeDomain(tt.args.domain); got != tt.want {
				t.Errorf("NormalizeDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateHostname(t *testing.T) {
	type args struct {
		hostname string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateHostname(tt.args.hostname); got != tt.want {
				t.Errorf("ValidateHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateDNSRecord(t *testing.T) {
	type args struct {
		hostname string
		ip       string
		TTL      int64
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
			if err := UpdateDNSRecord(tt.args.hostname, tt.args.ip, tt.args.TTL); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDNSRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateDNS(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateDNS(tt.args.w, tt.args.r)
		})
	}
}
