package main

import (
	"debug/dwarf"
	"time"
)

const timeOut = time.Second * 10

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	telnet := NewTelnetClient("127.0.01:3302", timeOut)

}
