package main

import (
	"aribtool/internal/dumpepg"
)

func main() {
	path := "/Users/swkoori/Downloads/7FE57FE5_epg.dat"
	//path := "/Users/swkoori/Documents/2021-0314-2148_東海テレビ０１１_東海テレニュース[字]【２０日（土）よる１１時４０分スタート！「リカ～リバース～」】.ts"

	dumpepg.DumpEpg(path, 1064)

}
