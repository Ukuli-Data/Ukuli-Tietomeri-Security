// Make a list of IP addresses banned by Fail2Ban just for lulz.package Security

package main

import (
	"os"
	"fmt"
)

var (
	jail bool = true
)

func jailthem(hacker string) (jail bool) {
	jail = true
	f, err := os.OpenFile("hackers.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Oh no: ", err)
		return jail
	}
	fmt.Fprintf(f,hacker + "\n")
	f.Close()
	fmt.Println(jail)
	return jail
}

func main() () {
	hacker := os.Args[1]
	jailthem(hacker)
}
