package main

import (
	"aribtool/tsparser"
	"encoding/json"
	"os"
)

func main() {
	//path := "/Users/swkoori/Downloads/7FE57FE5_epg.dat"
	path := "/Users/swkoori/Documents/2021-0314-2148_東海テレビ０１１_東海テレニュース[字]【２０日（土）よる１１時４０分スタート！「リカ～リバース～」】.ts"
	eitSectionList := tsparser.Scan(path, tsparser.ScheduleEventSection)

	f, _ := os.Create("data_output.json")
	defer f.Close()

	/*
		for _, v := range eitSectionList {
			fmt.Println(v)
		}*/
	events := tsparser.ParseEventSection(eitSectionList...)
	for _, e := range events {
		//jsonBytes, _ := json.Marshal(e)
		json.NewEncoder(f).Encode(e)
		//fmt.Println(string(json))
		//fmt.Printf("%+v\n", e)
	}

	/*
		for _, t := range tables {
			for _, e := range t.Event {
				if e.EventId == 4367 {
					if len(e.Descriptor.ExtendedEventDescriptor) > 0 {
						fmt.Printf("%+v\n", e)
						os.Exit(0)
					}
				}
			}
		}*/

}
