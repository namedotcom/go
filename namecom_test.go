package namecom

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var (
	namecom *NameCom
	once    sync.Once
)

func init() {
	once.Do(func() {
		namecom = New(os.Getenv("NAME_COM_USER"), os.Getenv("NAME_COM_TOKEN"))

		if os.Getenv("NAME_COM_PROXY") != "" {
			// setup a http client
			httpTransport := &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			}

			// create a socks5 dialer
			dialer, err := proxy.SOCKS5("tcp", os.Getenv("NAME_COM_PROXY"), nil, proxy.Direct)
			if err != nil {
				log.Fatal(err)
			}

			// set our socks5 as the dialer
			if contextDialer, ok := dialer.(proxy.ContextDialer); ok {
				httpTransport.DialContext = contextDialer.DialContext
			}

			namecom.Client = &http.Client{
				Transport: httpTransport,
				Timeout:   10 * time.Second,
			}
		} // proxy set for white ip list
	})
}
