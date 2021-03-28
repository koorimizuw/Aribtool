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
		for _, e := range v.Event {
			e.OriginalNetworkId = v.OriginalNetworkId
			e.TransportStreamId = v.TransportStreamId
			e.ServiceId = v.ServiceId
			events = append(events, e)
		}
	}
	return events
}
