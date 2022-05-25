package main

import (
	"fmt"
	"sshcrack/plugin"
)

func main() {
	fmt.Print("[+] ssh crack start")
	plugin.Checkinit()
	plugin.HandleCraklist()
}
