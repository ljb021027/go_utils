package x509

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	pemString := `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKhPSTDs4cpKfnMc
p86fCkpnuER7bGc+mGkhkw6bE+BnROfrDCFBSjrENLS5JcsenANQ1kYGt9iVW2fd
ZAWUdDoj+t7g6+fDpzY1BzPSUls421Dmu7joDPY8jSdMzFCeg7Lyj0I36bJJ7ooD
VPW6Q0XQcb8FfBiFPAKuY4elj/YDAgMBAAECgYBo2GMWmCmbM0aL/KjH/KiTawMN
nfkMY6DbtK9/5LjADHSPKAt5V8ueygSvI7rYSiwToLKqEptJztiO3gnls/GmFzj1
V/QEvFs6Ux3b0hD2SGpGy1m6NWWoAFlMISRkNiAxo+AMdCi4I1hpk4+bHr9VO2Bv
V0zKFxmgn1R8qAR+4QJBANqKxJ/qJ5+lyPuDYf5s+gkZWjCLTC7hPxIJQByDLICw
iEnqcn0n9Gslk5ngJIGQcKBXIp5i0jWSdKN/hLxwgHECQQDFKGmo8niLzEJ5sa1r
spww8Hc2aJM0pBwceshT8ZgVPnpgmITU1ENsKpJ+y1RTjZD6N0aj9gS9UB/UXdTr
HBezAkEAqkDRTYOtusH9AXQpM3zSjaQijw72Gs9/wx1RxOSsFtVwV6U97CLkV1S+
2HG1/vn3w/IeFiYGfZXLKFR/pA5BAQJAbFeu6IaGM9yFUzaOZDZ8mnAqMp349t6Q
DB5045xJxLLWsSpfJE2Y12H1qvO1XUzYNIgXq5ZQOHBFbYA6txBy/QJBAKDRQN47
6YClq9652X+1lYIY/h8MxKiXpVZVncXRgY6pbj4pmWEAM88jra9Wq6R77ocyECzi
XCqi18A/sl6ymWc=
-----END PRIVATE KEY-----`

	block, _ := pem.Decode([]byte(pemString))
	parseResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parseResult.(*rsa.PrivateKey)
	fmt.Println(key.N)
}

func main() {

}
