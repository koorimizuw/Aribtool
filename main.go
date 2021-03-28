package main

import (
	"aribtool/tsparser"
	"fmt"
)

func main() {
	var path string = "D:\\7FE27FE2_epg.dat"
	eitSectionList := tsparser.Scan(path, tsparser.EventSection)
	events := tsparser.ParseEventSection(eitSectionList...)
	/*for _, v := range events {
		fmt.Printf("%+v\n", v)
	}*/
	for i, v := range events {
		fmt.Printf("%+v\n", v)
		if i > 10 {
			break
		}
	}
}
