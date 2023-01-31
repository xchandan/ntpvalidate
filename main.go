package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func usage() {
	fmt.Printf("%s: <ntp server to validate>", os.Args[0])
}
func main() {
	fmt.Println("Validating ntp server")
	if len(os.Args) != 2 {
		usage()
	}
	ntpserver := os.Args[1]
	ntpresp, err := ntp.Query(ntpserver)
	if err != nil {
		fmt.Printf("Validation for ntp server %s failed with %s\n", ntpserver, err)
		os.Exit(1)
	}
	fmt.Printf("NTP server has offset of %s at stratum %d\n", ntpresp.ClockOffset, ntpresp.Stratum)
}
