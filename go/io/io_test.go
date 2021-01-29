package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	reader := bytes.NewReader(bs)

	now := time.Now()
	fmt.Println(now)
	_, err := ioutil.ReadAll(reader)
	fmt.Println(time.Since(now))
	assert.NoError(t, err)
}

func TestHttpLoc(t *testing.T) {
	get, err := http.DefaultClient.Get("http://localhost:28000")
	assert.NoError(t, err)
	reader := get.Body

	now := time.Now()
	fmt.Println(now)
	_, err = ioutil.ReadAll(reader)
	fmt.Println(time.Since(now))
}

func TestHttp(t *testing.T) {
	get, err := http.DefaultClient.Get("http://9.135.95.70:28000")
	assert.NoError(t, err)
	reader := get.Body

	now := time.Now()
	fmt.Println(now)
	_, err = ioutil.ReadAll(reader)
	fmt.Println(time.Since(now))
}

//var rspBody = []byte(`result_code=0&result_data=-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAyBytqTuj2m3BSnu3ISiX7Mfl6+EzRt10TXid1ClHsfjHe6Cr\nmB+pMQT/Wws8lNuqc0UrfjXQFt6Fd7AvTUr2eMSRGTXmnTICeyOg5ga6tEJDjwJQ\nLyHBFG00e56QBzCO3c56LiWBAwtTpq592FZQ9tcdhd+WnyNB6xa/C1XxnpG4fJcT\nPSFzmWwe5HktBAYap5MGCTmPJ57K+IO+GABnXRNeBsD+p0AjSQD1GqBGVwCdkOxr\nB5LllLpsdu2KGvTiRDORPwNL0iSrByxez9taFHum3h2+GSTK9+NdjPcEX2e/gg1E\nj1dj9BYU2nTp+ggGpnWxi9nnobnAin8Uj1BjJQIDAQABAoIBAHpStYE2eMGjx7z9\nBQOa6cuOoihfP9X4twhIS8oN9cp/tYkHvPy/lc/mKyX/J2gjIv1VcfGzeWERYuq3\n4yJPSXynQ77yaOb8U/Hr5IXX5TcdLmjz1AoLixgubR+H3KjYcOx7M8qTmFpBCUBh\nd4HPuaw58M0N0Oe0SK4o+F04ivsPKb+hVzAdXfwQjIq/FwTeucFZ78mVfnOw8ELw\n0was6HIyrAtMbPgKOHsTrYGLaKeXG7GdDiyr6YuDuXMxA4nhFXZJEe3E9zQ/tUHA\nqZ9tcF+UnuP/sjw9HVNwozKK66+g+3lec4TcIDHSwzt6Pku0JfG08WQSNJrmstrd\n30RGg90CgYEA5txP+rvxR98OMy+lDWDQFd1Vl+zBb+uZiSqSsXDFTbtlyeD7daAK\nFeN+TlwO178qEwvdTOhqjC+YF02ChDWlA0mLlFWIVqGPAyJLGVavFz3fMGm9XHoe\n6npmY3DOHsvHnx10a+mFjXmMxd9FIL8FK17RqmMUmXjCHB1RGnVdkesCgYEA3ecw\nlyTQQKC042JljFEd9jhACvgu/eBK1Z5Da4ay0HtMI3fnDoywh/DHT5e6iguwQVjD\nbs6LQzfYO0M7Uk4Bhjj21FxJzdmZ234PsBMyYlx/9Yvh6qxjnkvf3lY4ieYHqswY\nFvnuNCxDLNSEJS/GNXquMBa4iIPhK1jCEfy4iy8CgYEAmU6aqjYev4ynwGQWZnDx\nHmPyUEaAGpPJTOaBreXmkH0u1kZalr0llW2SVfUaQjmHdkh9uHHrN5bvmCH311ZM\nc4dmcqXuOSFmeD1Qw83lzjhfO6YsgvnyB+IqlkS2DItk0AxeYL8d2lpqmohvPpS7\nmRaJg23yowwnU0ZAsqJX9esCgYA3gx2RwiDi+hEmjTgQOT7AmOPUmq/OzSyLcjM1\nlBXpt96dROKlriZ/a4nA7Jk67Z+jSVxivQLzxuWuNayCc4dIF7oP+CJyf8xQr6Gn\nD7ZcCv3r5JmYxpWJRTv0+CbewJAMd1BTdyTTInuPnH6Oevwf5tfqqdRyOmO7H1I1\ng756+wKBgQDcs7Y4nDVgG9KOEAANqwG3DmfXRVBGUI3XBlPdalAJkKQqtfEjH+Vl\nD6lyLPemsi7knzN8U6pUX6CNspEqZ9KIkJInhyIpXZINAkRHZ2dm9KrkLipDTbdI\nPynUgiPGYx0tO+xoDpoRoR6R8FnUEOXa3lC85wMgcxPr7w5LMW+LAw==\n-----END RSA PRIVATE KEY-----&result_info=OK`)
var rspBody = []byte(`result_code=0&result_data=-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAyBytqTuj2m3BSnu3ISiX7Mfl6+EzRt10TXid1ClHsfjHe6Cr
mB+pMQT/Wws8lNuqc0UrfjXQFt6Fd7AvTUr2eMSRGTXmnTICeyOg5ga6tEJDjwJQ
LyHBFG00e56QBzCO3c56LiWBAwtTpq592FZQ9tcdhd+WnyNB6xa/C1XxnpG4fJcT
PSFzmWwe5HktBAYap5MGCTmPJ57K+IO+GABnXRNeBsD+p0AjSQD1GqBGVwCdkOxr
B5LllLpsdu2KGvTiRDORPwNL0iSrByxez9taFHum3h2+GSTK9+NdjPcEX2e/gg1E
j1dj9BYU2nTp+ggGpnWxi9nnobnAin8Uj1BjJQIDAQABAoIBAHpStYE2eMGjx7z9
BQOa6cuOoihfP9X4twhIS8oN9cp/tYkHvPy/lc/mKyX/J2gjIv1VcfGzeWERYuq3
4yJPSXynQ77yaOb8U/Hr5IXX5TcdLmjz1AoLixgubR+H3KjYcOx7M8qTmFpBCUBh
d4HPuaw58M0N0Oe0SK4o+F04ivsPKb+hVzAdXfwQjIq/FwTeucFZ78mVfnOw8ELw
0was6HIyrAtMbPgKOHsTrYGLaKeXG7GdDiyr6YuDuXMxA4nhFXZJEe3E9zQ/tUHA
qZ9tcF+UnuP/sjw9HVNwozKK66+g+3lec4TcIDHSwzt6Pku0JfG08WQSNJrmstrd
30RGg90CgYEA5txP+rvxR98OMy+lDWDQFd1Vl+zBb+uZiSqSsXDFTbtlyeD7daAK
FeN+TlwO178qEwvdTOhqjC+YF02ChDWlA0mLlFWIVqGPAyJLGVavFz3fMGm9XHoe
6npmY3DOHsvHnx10a+mFjXmMxd9FIL8FK17RqmMUmXjCHB1RGnVdkesCgYEA3ecw
lyTQQKC042JljFEd9jhACvgu/eBK1Z5Da4ay0HtMI3fnDoywh/DHT5e6iguwQVjD
bs6LQzfYO0M7Uk4Bhjj21FxJzdmZ234PsBMyYlx/9Yvh6qxjnkvf3lY4ieYHqswY
FvnuNCxDLNSEJS/GNXquMBa4iIPhK1jCEfy4iy8CgYEAmU6aqjYev4ynwGQWZnDx
HmPyUEaAGpPJTOaBreXmkH0u1kZalr0llW2SVfUaQjmHdkh9uHHrN5bvmCH311ZM
c4dmcqXuOSFmeD1Qw83lzjhfO6YsgvnyB+IqlkS2DItk0AxeYL8d2lpqmohvPpS7
mRaJg23yowwnU0ZAsqJX9esCgYA3gx2RwiDi+hEmjTgQOT7AmOPUmq/OzSyLcjM1
lBXpt96dROKlriZ/a4nA7Jk67Z+jSVxivQLzxuWuNayCc4dIF7oP+CJyf8xQr6Gn
D7ZcCv3r5JmYxpWJRTv0+CbewJAMd1BTdyTTInuPnH6Oevwf5tfqqdRyOmO7H1I1
g756+wKBgQDcs7Y4nDVgG9KOEAANqwG3DmfXRVBGUI3XBlPdalAJkKQqtfEjH+Vl
D6lyLPemsi7knzN8U6pUX6CNspEqZ9KIkJInhyIpXZINAkRHZ2dm9KrkLipDTbdI
PynUgiPGYx0tO+xoDpoRoR6R8FnUEOXa3lC85wMgcxPr7w5LMW+LAw==
-----END RSA PRIVATE KEY-----
&result_info=OK`)

func TestHttpService(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(rspBody)
	})
	http.ListenAndServe(":28000", nil)

}
