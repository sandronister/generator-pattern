package request

import (
	"io"
	"net/http"
	"regexp"
	"time"
)

func GetTitle(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			html, err := io.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}

			r, _ := regexp.Compile("<title>(.*?)</title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)

	}
	return c
}

func GetBest(url1, url2, url3, url4 string) string {
	c1 := GetTitle(url1)
	c2 := GetTitle(url2)
	c3 := GetTitle(url3)
	c4 := GetTitle(url4)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case t4 := <-c4:
		return t4
	case <-time.After(1000 * time.Millisecond):
		return "Timeout"

	}
}
