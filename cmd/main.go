package main

import (
	"fmt"

	"github.com/sandronister/generator-pattern/pkg/mult"
	"github.com/sandronister/generator-pattern/pkg/request"
)

func main() {
	c := mult.Combine(
		request.GetTitle("https://www.google.com", "https://www.bing.com"),
		request.GetTitle("https://www.apple.com/br/", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
