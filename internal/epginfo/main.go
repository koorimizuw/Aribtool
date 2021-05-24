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

func EpgInfo(path string) (eventData *EventData) {
	eventPid := tsparser.PidMap[tsparser.CurrentEventSection]
	eventTidRange := tsparser.TableIdMap[tsparser.CurrentEventSection]

	eventSectionList := tsparser.Scan(path, eventPid, eventTidRange, 100)
	events := tsparser.ParseCurrentEventSection(81, eventSectionList...)
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
