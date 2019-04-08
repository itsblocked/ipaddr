// ipaddr is a simple library to validate an IPv4 internet protocol (ip) address
// with a port.
package ipaddr

import (
	"fmt"
	"strconv"
	"strings"
)

// IpAddr is a string IP address of the form "###.###.###.###:port".
type IpAddr string

// String is a "stringer" for IpAddr.
func (ip *IpAddr) String() string {
	return string(*ip)
}

// Set assigns value to ip if it validates.
func (ip *IpAddr) Set(value string) error {
	parts := strings.Split(value, ":")
	if len(parts) != 2 || len(parts[1]) == 0 {
		return fmt.Errorf("port # is required")
	}

	ipParts := strings.Split(parts[0], ".")
	if len(ipParts) != 4 {
		return fmt.Errorf("invalid tcp4 IP address")
	}

	for _, x := range ipParts {
		if i, err := strconv.Atoi(x); err != nil {
			return err
		} else {
			if i < 0 || i > 255 {
				return fmt.Errorf("invalid number (%s) in IP address", x)
			}
		}
	}

	*ip = IpAddr(value)

	return nil
}

// Port returns the port portion of IpAddr or empty string if port is not found.
func (ip *IpAddr) Port() string {
	parts := strings.Split(string(*ip), ":")
	if len(parts) != 2 || len(parts[1]) == 0 {
		return ""
	}
	if _, err := strconv.Atoi(parts[1]); err != nil {
		return ""
	}
	return parts[1]
}

// vim: sw=8 sts=8 noet
