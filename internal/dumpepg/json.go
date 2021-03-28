package dumpepg

import (
	"encoding/json"
	"fmt"
	"os"
)

func dump(data []EventData) {
	filename := fmt.Sprintf("%d-%d-%d", data[0].Onid, data[0].Tsid, data[0].Sid)
	f, _ := os.Create(filename + ".json")
	defer f.Close()

	//fmt.Println(data)
	json.NewEncoder(f).Encode(data)
}
