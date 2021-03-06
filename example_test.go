package netutil_test

import (
	"fmt"
	"testing"

	"github.com/HankB/netutil"
)

func TestNetUtil(t *testing.T) {
	// test both IPV$ and IPV6
	bcast, err := netutil.GetBroadcastIPs(true)
	if err != nil {
		t.Fail()
	}
	if len(bcast) < 2 {
		fmt.Println("less than 2 broadcast addresses found", len(bcast))
		t.Fail()
	}
	fmt.Println(err, bcast)

	// test IPV4 only
	bcast, err = netutil.GetBroadcastIPs(false)
	if err != nil {
		t.Fail()
	}
	if len(bcast) < 1 {
		fmt.Println("No IPV4 broadcast address found", len(bcast))
		t.Fail()
	}
	fmt.Println(err, bcast)

}
