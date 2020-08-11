// Make a list of IP addresses banned by Fail2Ban just for lulz.package Security

package main

import (
	"os"
	"fmt"
)

func main() {
	hacker := os.Args[1]
	f, err := os.OpenFile("hackers.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Oh no: ", err)
		return
	}
	fmt.Fprintf(f,hacker + "\n")
	f.Close()
	return
}
