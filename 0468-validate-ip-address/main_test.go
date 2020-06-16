package main

import "testing"

var tests = []struct {
	IP   string
	want string
}{
	// {"172.16.254.1", "IPv4"},
	// {"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
	{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "IPv6"},
	// {"256.256.256.256", "Neither"},
	// {"2001:0db8:85a3::8A2E:0370:7334", "Neither"},
	// {"02001:0db8:85a3:0000:0000:8a2e:0370:7334", "Neither"},
	// {"2001:0dx8:85a3:0000:0000:8a2e:0370:7334", "Neither"},
}

func TestValidIPAddress(t *testing.T) {
	for _, test := range tests {
		if got := validIPAddress(test.IP); got != test.want {
			t.Errorf("validIPAddress(%q) = %q; want %q\n", test.IP, got, test.want)
		}
	}
}
