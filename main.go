package main

import (
	"flag"
	"fmt"
	"net/http"
)

var host = flag.String("host", "https://example.com", "Host to connect to.")
var attempts = flag.Int("attempts", 100, "Number of attempts to make.")

func main() {
	flag.Parse()

	fmt.Printf("Will attempt to connect %d times to %s\n", *attempts, *host)
	for i := 0; i < *attempts; i++ {
		connectAndTest()
	}
}

func connectAndTest() {
	_, err := http.Get(*host)
	if err != nil {
		fmt.Println(err)
	}
}
