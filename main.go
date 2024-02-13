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

	fmt.Println("Test complete.")
}

func connectAndTest() {
	req, err := http.NewRequest("GET", *host, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "https://github.com/aaomidi/connection_test")

	if _, err := http.DefaultClient.Do(req); err != nil {
		fmt.Println(err)
	}
}
