package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"lenovo/handle"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	kw := flag.String("kw", "test", "key word")
	proxy := flag.String("proxy", "null", "proxy")
	flag.Parse()
	pvid := handle.PVid()
	client := http.Client{}
	if *proxy != "null" {
		fmt.Println("设置代理为:", *proxy)
		transport := &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse("http://" + *proxy)
		}}
		client = http.Client{Transport: transport}
	}

	req, err2 := http.NewRequest("GET", handle.HandlerUrl(*kw, pvid), strings.NewReader(""))
	if err2 != nil {
		panic(err2)
	}
	req.Header.Set("user-agent", RandomString())
	req.Header.Set("Referer", handle.HandlerReferer())
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//line, _, err := bufio.NewReader(res.Body).ReadLine()
	fmt.Println(string(all))
}
