package request

import (
	"io"
	"net/http"
	"regexp"
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
