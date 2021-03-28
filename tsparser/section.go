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

func ParseEventSection(sid int, sectionList ...[]byte) (int, int, []table.Event) {
	var onid int
	var tsid int
	var events []table.Event
	for i, v := range table.ParseEventSection(sectionList...) {
		if i == 0 {
			onid = v.OriginalNetworkId
			tsid = v.TransportStreamId
		}
		if v.ServiceId != sid {
			continue
		}
		events = append(events, v.Event...)
	}
	return onid, tsid, events
}
