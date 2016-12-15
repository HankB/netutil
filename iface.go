package netutil

import (
	"net"
	"strings"
)

// GetBroadcastIPs identifies broadcast addresses on a given host ignoring
// local interfaces
func GetBroadcastIPs(includeIPV6 bool) ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	var bcastAddrs []net.IP
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if strings.Compare(addr.String(), "127.0.0.1/8") == 0 ||
			strings.Compare(addr.String(), "::1/128") == 0 {
		} else {
			addrIP, ipNet, _ := net.ParseCIDR(addr.String())
			bcast := ipNet.IP
			for i := range ipNet.Mask {
				bcast[i] |= ^ipNet.Mask[i]
			}
			if addrIP.DefaultMask() != nil {
				bcastAddrs = append(bcastAddrs, bcast)
			} else if includeIPV6 {
				bcastAddrs = append(bcastAddrs, bcast)
			}
		}
	}
	return bcastAddrs, nil
}
