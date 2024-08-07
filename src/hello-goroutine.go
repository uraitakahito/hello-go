//go:build ignore

// https://learn.microsoft.com/ja-jp/training/modules/go-concurrency/1-goroutines
package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string) {
    _, err := http.Get(api)
    if err != nil {
        fmt.Printf("ERROR: %s is down!\n", api)
        return
    }

    fmt.Printf("SUCCESS: %s is up and running!\n", api)
}
func main() {
    start := time.Now()

    apis := []string{
        "https://management.azure.com",
        "https://dev.azure.com",
        "https://api.github.com",
        "https://outlook.office.com/",
        "https://api.somewhereintheinternet.com/",
        "https://graph.microsoft.com",
    }

		for _, api := range apis {
				go checkAPI(api)
		}

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
