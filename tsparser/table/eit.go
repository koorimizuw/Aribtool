package table

import (
	"time"
)

/*
event_information_section(){
		table_id 						8 uimsbf
		section_syntax_indicator 		1 bslbf
		reserved_future_use 			1 bslbf
		reserved 						2 bslbf
		section_length 					12
		service_id 						16 uimsbf
		reserved 						2 bslbf
		version_number 					5 uimsbf
		current_next_indicator 			1 bslbf
		section_number 					8 uimsbf
		last_section_number 			8 uimsbf
		transport_stream_id 			16 uimsbf
		original_network_id 			16 uimsbf
		segment_last_section_number 	8 uimsbf
		last_table_id 					8 uimsbf
		for(i=0;i<N;i++){
			event_id 					16 uimsbf
			start_time 					40 bslbf
			duration 					24 uimsbf
			running_status 				3 uimsbf
			free_CA_mode 				1 bslbf
			descriptors_loop_length 	12 uimsbf
			for(i=0;i<N;i++){
			descriptor()
		}
	}
	CRC_32 								32 rpchof
}*/
type EventSection struct {
	TableId                  int
	SectionSyntaxIndicator   bool
	SectionLength            int
	ServiceId                int
	VersionNumber            int
	CurrentNextIndicator     bool
	SectionNumber            int
	LastSectionNumber        int
	TransportStreamId        int
	OriginalNetworkId        int
	SegmentLastSectionNumber int
	LastTableId              int
	Event                    []Event
	CRC32                    []byte
}

type Event struct {
	OriginalNetworkId     int
	TransportStreamId     int
	ServiceId             int
	EventId               int
	StartTime             time.Time
	Duration              int
	RunningStatus         int
	FreeCAMode            bool
	DescriptorsLoopLength int
	Descriptor            Descriptor
}

func ParseEventSection(sectionList ...[]byte) []EventSection {
	var s Section
	var tmp []EventSection
	for _, section := range sectionList {
		s = section
		eventSection := EventSection{
			TableId:                  s.uimsbf(0, 8).toNumber(),
			SectionSyntaxIndicator:   s.bslbf(8, 1).toBool(),
			SectionLength:            s.bslbf(12, 12).toNumber(),
			ServiceId:                s.uimsbf(24, 16).toNumber(),
			VersionNumber:            s.uimsbf(42, 5).toNumber(),
			CurrentNextIndicator:     s.bslbf(47, 1).toBool(),
			SectionNumber:            s.uimsbf(48, 8).toNumber(),
			LastSectionNumber:        s.uimsbf(56, 8).toNumber(),
			TransportStreamId:        s.uimsbf(64, 16).toNumber(),
			OriginalNetworkId:        s.uimsbf(80, 16).toNumber(),
			SegmentLastSectionNumber: s.uimsbf(96, 8).toNumber(),
			LastTableId:              s.uimsbf(104, 8).toNumber(),
			Event:                    parseEvent(s[14 : len(s)-4]),
			CRC32:                    s[len(s)-4:],
		}
		tmp = append(tmp, eventSection)
	}

	return tmp
}

func parseEvent(s Section) []Event {
	var tmp []Event
	for len(s) > 0 {
		descriptorsLoopLength := s.uimsbf(84, 12).toNumber()
		tmp = append(tmp, Event{
			EventId:               s.uimsbf(0, 16).toNumber(),
			StartTime:             s.bslbf(16, 40).toTime(),
			Duration:              s.uimsbf(56, 24).toDuration(),
			RunningStatus:         s.uimsbf(80, 3).toNumber(),
			FreeCAMode:            s.bslbf(83, 1).toBool(),
			DescriptorsLoopLength: descriptorsLoopLength,
			Descriptor:            parseDescriptor(s[12 : 12+descriptorsLoopLength]),
		})
		s = s[12+descriptorsLoopLength:]
	}
	return tmp
}

func (s Section) IsEventSection() bool {
	tid := s.TableId()
	if tid >= 0x4E && tid <= 0x6F {
		return true
	}
	return false
}
