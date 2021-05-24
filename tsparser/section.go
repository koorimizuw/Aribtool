package tsparser

import (
	"github.com/koorimizuw/aribtool/tsparser/table"
)

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

func ParseCurrentEventSection(skip int, sectionList ...[]byte) []table.Event {
	var events []table.Event
	skipCount := 0
	for _, v := range table.ParseEventSection(sectionList...) {
		if skipCount > skip && v.SectionNumber == 0 {
			for _, e := range v.Event {
				e.OriginalNetworkId = v.OriginalNetworkId
				e.TransportStreamId = v.TransportStreamId
				e.ServiceId = v.ServiceId
				events = append(events, e)
			}
		}
		skipCount += 1
	}
	return events
}

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
