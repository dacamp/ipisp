package ipisp

import (
	"reflect"
	"testing"
)

func TestParseASN(t *testing.T) {
	type args struct {
		asn string
	}
	tests := []struct {
		name    string
		args    args
		want    ASN
		wantErr bool
	}{
		{"", args{"AS555"}, ASN(555), false},
		{"", args{"AS"}, ASN(0), true},
		{"", args{"ASDFASDF"}, ASN(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseASN(tt.args.asn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseASN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseASN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestASN_String(t *testing.T) {
	tests := []struct {
		name string
		a    ASN
		want string
	}{
		{"", ASN(555), "AS555"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("ASN.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASNs(t *testing.T) {
	type args struct {
		asnList string
	}
	tests := []struct {
		name    string
		args    args
		want    []ASN
		wantErr bool
	}{
		{"single", args{"1010"}, []ASN{1010}, false},
		{"double", args{"1010 1000"}, []ASN{1010, 1000}, false},
		{"none", args{""}, nil, true},
		{"bad", args{"hello"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseASNs(tt.args.asnList)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseASNs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASNs() = %v, want %v", got, tt.want)
			}
		})
	}
}
