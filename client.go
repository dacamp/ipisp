package ipisp

import (
	"net"
)

// Client represents an IP or ASN lookup client.
type Client interface {
	// LookupIPs looks up IPs and returns a slice of responses the same size as the input slice of IPs
	// in the same order.
	LookupIPs([]net.IP) ([]Response, error)
	LookupIP(net.IP) (*Response, error)
	LookupASNs([]ASN) ([]Response, error)
	LookupASN(ASN) (*Response, error)
	Close() error
}
