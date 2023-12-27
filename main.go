package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func getTitle(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)</title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)

	}
	return c
}

func main() {
	t1 := getTitle("https://www.google.com", "https://www.bing.com")
	t2 := getTitle("https://www.amazon.com.br", "https://www.youtube.com")

	fmt.Println("Primeiro", <-t1, "|", <-t2)
	fmt.Println("Segundo", <-t1, "|", <-t2)
}
