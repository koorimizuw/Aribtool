package tsparser

import "aribtool/tsparser/table"

type TableIdRange struct {
	Start byte
	End   byte
}

var TableIdMap = map[string]TableIdRange{
	CurrentEventSection:  {0x4E, 0x4E},
	ScheduleEventSection: {0x50, 0x6F},
	// ...
}

const (
	EventSection         = "EventSection"
	CurrentEventSection  = "CurrentEventSection"
	ScheduleEventSection = "ScheduleEventSection"
)

func ParseEventSection(sectionList ...[]byte) []table.Event {
	var events []table.Event
	for _, v := range table.ParseEventSection(sectionList...) {
		events = append(events, v.Event...)
	}
	return events
}
