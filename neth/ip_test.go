package neth

import (
	"net"
	"testing"
)

var tests = []struct {
	ip     net.IP
	isV4   bool
	isV6   bool
	length int
}{
	{isV4: false, isV6: false, length: 0},
	{ip: []byte{1, 2}, isV4: false, isV6: false, length: 0},
	{ip: []byte{1, 2, 3, 4}, isV4: true, isV6: false, length: 4},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3}, isV4: false, isV6: false, length: 0},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4}, isV4: true, isV6: false, length: 4},
	{ip: []byte{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4}, isV4: false, isV6: true, length: 16},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4, 5}, isV4: false, isV6: false, length: 0},
}

func TestIsIPv4(t *testing.T) {
	for _, test := range tests {
		if r := IsIPv4(test.ip); r != test.isV4 {
			t.Errorf("%v: expect %v, got %v", test.ip, test.isV4, r)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	for _, test := range tests {
		if r := IsIPv6(test.ip); r != test.isV6 {
			t.Errorf("%v: expect %v, got %v", test.ip, test.isV6, r)
		}
	}
}

func TestIPLen(t *testing.T) {
	for _, test := range tests {
		if r := IPLen(test.ip); r != test.length {
			t.Errorf("%v: expect %v, got %v", test.ip, test.length, r)
		}
	}
}
