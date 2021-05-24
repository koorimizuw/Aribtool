package main

import (
	"aribtool/internal/dumpepg"
	"aribtool/internal/epginfo"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	printDefault := func() {
		fmt.Print("Aribtool")
	}

	if len(args) == 0 {
		printDefault()
		return
	}

	switch args[0] {
	case "dumpepg":
		if len(args) == 1 {
			printDefault()
		}
		dumpepg.DumpEpg(args[1])
	case "epginfo":
		if len(args) == 1 {
			printDefault()
		}
		epginfo.EpgInfo(args[1])
	default:
		printDefault()
	}
}
