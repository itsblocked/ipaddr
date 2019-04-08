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

func TestIpAddrPort(t *testing.T) {
	cases := map[string]string{
		"127.0.0.1:4040":  "4040",
		"10.1.10.1:9000":  "9000",
		"127.0.0.1:55555": "55555",
		"127.0.0.1:9":     "9",
	}

	for s, want := range cases {
		var ip IpAddr
		err := ip.Set(s)
		if err != nil {
			t.Fatal(err)
		}

		p := ip.Port()
		if p != want {
			t.Errorf("got: %s, want %s", p, want)
		}
	}
}

// vim: sw=8 sts=8 noet
