package tsparser

import "aribtool/tsparser/table"

type TableIdRange struct {
	Start byte
	End   byte
}

var TableIdMap = map[string]TableIdRange{
	EventSection: {0x4E, 0x6F},
}

const (
	EventSection    = "EventSection"
	EventSectionEx1 = "EventSectionEx1"
	EventSectionEx2 = "EventSectionEx2"
)

func ParseEventSection(sectionList ...[]byte) []table.EventSection {
	return table.ParseEventSection(sectionList...)
}
