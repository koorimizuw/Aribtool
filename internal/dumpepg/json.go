package dumpepg

import (
	"encoding/json"
	"fmt"
	"os"
)

func dump(data []EventData) {
	if len(data) == 0 {
		panic("No event data.")
	}

	filename := fmt.Sprintf("%d", data[0].Onid)
	f, _ := os.Create(filename + ".json")
	defer f.Close()

	json.NewEncoder(f).Encode(data)
}
