package epginfo

import (
	"github.com/koorimizuw/aribtool/tsparser"
)

type EventData struct {
	Onid      int    `json:"onid"`
	Tsid      int    `json:"tsid"`
	Sid       int    `json:"sid"`
	Eid       int    `json:"eid"`
	StartTime string `json:"start_time"`
	Duration  int    `json:"duration"`
	EventName string `json:"event_name"`
}

func EpgInfo(path string, read int, skip int) (eventData *EventData) {
	patPid := tsparser.PidMap[tsparser.ProgramAssociationSection]
	patTidRange := tsparser.TableIdMap[tsparser.ProgramAssociationSection]
	patSectionList := tsparser.Scan(path, patPid, patTidRange, 100)
	sid := tsparser.GetSid(patSectionList[0])

	eventPid := tsparser.PidMap[tsparser.CurrentEventSection]
	eventTidRange := tsparser.TableIdMap[tsparser.CurrentEventSection]

	eventSectionList := tsparser.Scan(path, eventPid, eventTidRange, read)
	events := tsparser.ParseCurrentEventSection(sid, skip, eventSectionList...)
	event := events[len(events)-1]

	return &EventData{
		Onid:      event.OriginalNetworkId,
		Tsid:      event.TransportStreamId,
		Sid:       event.ServiceId,
		Eid:       event.EventId,
		StartTime: event.StartTime.Format("2006/01/02 15:04:05"),
		Duration:  event.Duration,
		EventName: event.Descriptor.ShortEventDescriptor[0].EventNameChar,
	}
}
