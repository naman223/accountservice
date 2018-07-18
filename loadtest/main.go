package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"strings"
)

var Log = logrus.New()

var baseAddr string
var zuul bool

func main() {

	usersPtr := flag.Int("users", 10, "Number of users")
	delayPtr := flag.Int("delay", 1000, "Delay between calls per user")
	zuulPtr := flag.Bool("zuul", true, "Route traffic through zuul")
	baseAddrPtr := flag.String("baseAddr", "localhost", "Base address of your Swarm cluster")

	flag.Parse()

	baseAddr = *baseAddrPtr
	zuul = *zuulPtr
	users := *usersPtr
	var _ int = *delayPtr

	for i := 0; i < users; i++ {
		// go securedTest()
		go standardTest()
	}

	// Block...
	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()

}

var defaultTransport http.RoundTripper = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

func standardTest() {
	var url string
	if zuul {
		Log.Println("Using HTTPS through ZUUL")
		url = "https://" + baseAddr + ":8765/api/accounts/"
	} else {
		url = "http://" + baseAddr + ":6767/accounts/"
	}
	m := make(map[string]interface{})
	for {
		accountId := rand.Intn(99) + 10000
		serviceUrl := url + strconv.Itoa(accountId)

		var DefaultTransport http.RoundTripper = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: false,
		}
		req, _ := http.NewRequest("GET", serviceUrl, nil)
		resp, err := DefaultTransport.RoundTrip(req)

		if err != nil {
			panic(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &m)
		quote :=  m["quote"].(map[string]interface{})["quote"].(string)
		quoteIp := m["quote"].(map[string]interface{})["ipAddress"].(string)
		quoteIp = quoteIp[strings.IndexRune(quoteIp, '/') + 1 :]

		imageUrl := m["imageData"].(map[string]interface{})["url"].(string)
		imageServedBy := m["imageData"].(map[string]interface{})["servedBy"].(string)

		fmt.Print("|" + m["name"].(string) + "\t|" + m["servedBy"].(string) + "\t|")
		fmt.Print(PadRight(quote, " ", 32) + "\t|" + quoteIp + "\t|")
		fmt.Println(PadRight(imageUrl, " ", 28) + "\t|" + imageServedBy + "\t|")
		time.Sleep(time.Second * 1)
	}

}

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
