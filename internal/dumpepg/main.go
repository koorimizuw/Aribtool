package dumpepg

import (
	"aribtool/tsparser"
	"aribtool/tsparser/table"
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

func DumpEpg(path string) {
	eventPid := tsparser.PidMap[tsparser.ScheduleEventSection]
	eventTidRange := tsparser.TableIdMap[tsparser.ScheduleEventSection]

	eventSectionList := tsparser.Scan(path, eventPid, eventTidRange)
	events := tsparser.ParseEventSection(eventSectionList...)

	var eventData []EventData
	eventMap := make(map[int]*EventData)
	for _, v := range events {
		if _, ok := eventMap[v.EventId]; !ok {
			eventMap[v.EventId] = &EventData{
				Onid:      v.OriginalNetworkId,
				Tsid:      v.TransportStreamId,
				Sid:       v.ServiceId,
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
			var lastDesc string
			desc := make(map[string]*[]byte)
			for _, ext := range v.Descriptor.ExtendedEventDescriptor {
				if ext.EventItem[0].ItemDescriptionLength != 0 {
					desc[ext.EventItem[0].ItemDescriptionChar] = &ext.EventItem[0].ItemChar
					lastDesc = ext.EventItem[0].ItemDescriptionChar
				} else {
					tmp := append(*desc[lastDesc], ext.EventItem[0].ItemChar...)
					desc[lastDesc] = &tmp
				}
			}

			for k, v := range desc {
				extendInfo += k + "\n"
				extendInfo += table.Mnemonic(*v).ToString() + "\n\n"
			}
			eventMap[v.EventId].EventDetailExt = extendInfo
		}
	}

	for _, v := range eventMap {
		eventData = append(eventData, *v)
	}

	dump(eventData)
}
