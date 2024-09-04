//go:build ignore

package main

// 1.
// The method by which Datadog-agent retrieves the hostname
// is documented in the following URL:
// https://github.com/DataDog/datadog-agent/blob/6f76f1ed2ef850c09e9130a094e4aae58ab2e5cd/pkg/util/hostname/providers.go#L53-L123
//
// 2.
// When `osHostnameUsable` returns true
// https://github.com/DataDog/datadog-agent/blob/6f76f1ed2ef850c09e9130a094e4aae58ab2e5cd/pkg/util/hostname/os_hostname_linux.go#L21-L41
//
// 3.
// When retrieving from the OS, the `fromOS` function is called.
// https://github.com/DataDog/datadog-agent/blob/6f76f1ed2ef850c09e9130a094e4aae58ab2e5cd/pkg/util/hostname/common.go#L104-L112
//
// 4.
// Finally, `os.Hostname` is called.

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Hostname Error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(hostname)
}
