package mult

func Foward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func Combine(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go Foward(input1, c)
	go Foward(input2, c)
	return c
}
