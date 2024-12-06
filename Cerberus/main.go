package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	ports  = make(chan int, 400)
	result = make(chan string, 400)
)

func main() {
	target := flag.String("t", "", "The target IP address.")
	ports := flag.String("p", "", "The port(s) to scan.")
	PandoraScan := flag.Bool("v", false, "Scan for pandora.")

	flag.Parse()
	if *target == "" || *ports == "" {
		fmt.Println("Error: Missing required flags.")
		flag.Usage()
		os.Exit(1)
	}
	ports_range := getPortsRange(*ports)
	ts := strings.Split(*target, ",")
	for _, target := range ts {
		scan(target, ports_range, *PandoraScan)

	}
}
