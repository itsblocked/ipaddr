package ipaddr

import (
	"testing"
)

func TestIpAddrSet(t *testing.T) {
	cases := map[string]bool{
		"127.0.0.1:4040": true,
		"10.1.10.1:9000": true,
		"127.0.0:4040":   false,
		"127.0.0.1,4040": false,
		"127.0.0.1:":     false,
	}

	for s, want := range cases {
		var ip IpAddr
		err := ip.Set(s)
		if err != nil {
			if want {
				t.Fatal(err)
			}
		} else if !want {
			t.Errorf("%s: expect fail", s)
		}
	}
}

// vim: sw=8 sts=8 noet
