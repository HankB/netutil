# netutil

Provide non-specific network utility code for various Go projects.

## Overview

### GetBroadcastIPs

Does what it says.

## Testing

`go test -v example_test.go`
should produce something that looks like
```
=== RUN   TestNetUtil
<nil> [192.168.1.255 2601:249:e00:b215:ffff:ffff:ffff:ffff 2601:249:e00:b215::14c2 2601:249:e00:b215:ffff:ffff:ffff:ffff 2601:249:e00:b215:ffff:ffff:ffff:ffff fe80::ffff:ffff:ffff:ffff]
<nil> [192.168.1.255]
--- PASS: TestNetUtil (0.00s)
PASS
ok  	command-line-arguments	0.002s
```