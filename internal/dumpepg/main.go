package dumpepg

import (
	"aribtool/tsparser"
)

type EventData struct {
	Onid           int    `json:"onid"`
	Tsid           int    `json:"tsid"`
	Sid            int    `json:"sid"`
	Eid            int    `json:"eid"`
	StartTime      string `json:"start_time"`
	Duration       int    `json:"duration"`
	EventName      string `json:"event_name"`
	EventDetail    string `json:"event_detail"`
	EventDetailExt string `json:"event_detail_ext"`
	Genre          int    `json:"ganre"`
}

func DumpEpg(path string, sid int) {
	eventSectionList := tsparser.Scan(path, tsparser.ScheduleEventSection)
	onid, tsid, events := tsparser.ParseEventSection(sid, eventSectionList...)

	eventMap := make(map[int]*EventData)
	for _, v := range events {
		if _, ok := eventMap[v.EventId]; !ok {
			eventMap[v.EventId] = &EventData{
				Onid:      onid,
				Tsid:      tsid,
				Sid:       sid,
				Eid:       v.EventId,
				StartTime: v.StartTime.Format("2006/01/02 15:04:05"),
				Duration:  v.Duration,
			}
		}

		if len(eventMap[v.EventId].EventName) == 0 && len(v.Descriptor.ShortEventDescriptor) > 0 {
			eventMap[v.EventId].EventName = v.Descriptor.ShortEventDescriptor[0].EventNameChar
			eventMap[v.EventId].EventDetail = v.Descriptor.ShortEventDescriptor[0].TextChar
		}

		if len(eventMap[v.EventId].EventDetailExt) == 0 && len(v.Descriptor.ExtendedEventDescriptor) > 0 {
			var extendInfo string
			for _, ext := range v.Descriptor.ExtendedEventDescriptor {
				extendInfo += ext.EventItem[0].ItemDescriptionChar
				extendInfo += "\n"
				extendInfo += ext.EventItem[0].ItemChar
				extendInfo += "\n\n"
			}
			eventMap[v.EventId].EventDetailExt = extendInfo
		}
	}

	var eventData []EventData
	for _, v := range eventMap {
		eventData = append(eventData, *v)
	}

	dump(eventData)
}
