package main

import (
	"flag"
	"fmt"

	"github.com/koorimizuw/aribtool/internal/dumpepg"
	"github.com/koorimizuw/aribtool/internal/epginfo"
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
		e := epginfo.EpgInfo(args[1], 200, 140)
		fmt.Printf("%#v", e)
	default:
		printDefault()
	}
}
