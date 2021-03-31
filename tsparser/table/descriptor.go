package table

import "fmt"

type Descriptor struct {
	ShortEventDescriptor     []ShortEventDescriptor     // 短形式イベント記述子
	ExtendedEventDescriptor  []ExtendedEventDescriptor  // 拡張形式イベント記述子
	ContentDescriptor        []ContentDescriptor        //コンテント記述子
	DataContentDescriptor    []DataContentDescriptor    // データコンテンツ記述子
	EventGroupDescriptor     []EventGroupDescriptor     //イベントグループ記述子
	ComponentDescriptor      []ComponentDescriptor      // コンポーネント記述子
	AudioComponentDescriptor []AudioComponentDescriptor // 音声コンポーネント記述子
}

func parseDescriptor(s Section) Descriptor {
	desc := Descriptor{}
	for len(s) > 0 {
		tag := s[0]
		length := int(s[1])

		payload := s[:2+length]

		switch tag {
		case 0x4d: // 短形式イベント記述子
			desc.ShortEventDescriptor = append(desc.ShortEventDescriptor, parseShortEventDescriptor(payload))
		case 0x4e: // 拡張形式イベント記述子
			desc.ExtendedEventDescriptor = append(desc.ExtendedEventDescriptor, parseExtendedEventDescriptor(payload))
		case 0x50: // コンポーネント記述子
			desc.ComponentDescriptor = append(desc.ComponentDescriptor, parseComponentDescriptor(payload))
		case 0x54: // コンテント記述子
			desc.ContentDescriptor = append(desc.ContentDescriptor, parseContentDescriptor(payload))
		case 0xc4: // 音声コンポーネント記述子
			desc.AudioComponentDescriptor = append(desc.AudioComponentDescriptor, parseAudioComponentDescriptor(payload))
		case 0xc1: // デジタルコピー制御記述子
			// 実装待ち
		case 0xc7: // データコンテンツ記述子
			desc.DataContentDescriptor = append(desc.DataContentDescriptor, parseDataContentDescriptor(payload))
		case 0xd6: //イベントグループ記述子
			// 実装待ち
		default:
			fmt.Println(tag)
		}
		s = s[2+length:]
	}

	return desc
}

/*
ConditionalAccessDescriptor
Copyright Descriptor
Network Name Descriptor
Service List Descriptor
Stuffing Descriptor
Satellite Delivery System Descriptor
Terrestrial Delivery System Descriptor
Bouquet Name Descriptor
Service Descriptor
Country Availability Descriptor
Linkage Descriptor
NVOD Reference Descriptor
Time Shifted Service Descriptor
*/

// 短形式イベント記述子
type ShortEventDescriptor struct {
	DescriptorTag      byte
	DescriptorLength   int
	ISO639LanguageCode string
	EventNameLength    int
	EventNameChar      string
	TextLength         int
	TextChar           string
}

func parseShortEventDescriptor(s Section) ShortEventDescriptor {
	eventNameLength := s.uimsbf(40, 8).toNumber()
	textLength := s.uimsbf(48+eventNameLength*8, 8).toNumber()
	return ShortEventDescriptor{
		DescriptorTag:      s.uimsbf(0, 8).toByte(),
		DescriptorLength:   s.uimsbf(8, 8).toNumber(),
		ISO639LanguageCode: s.bslbf(16, 24).toLanguageCode(),
		EventNameLength:    eventNameLength,
		EventNameChar:      s.uimsbf(48, eventNameLength*8).ToString(),
		TextLength:         textLength,
		TextChar:           s.uimsbf(56+eventNameLength*8, textLength*8).ToString(),
	}
}

// 拡張形式イベント記述子
type ExtendedEventDescriptor struct {
	DescriptorTag        byte
	DescriptorLength     int
	DescriptorNumber     int
	LastDescriptorNumber int
	ISO639LanguageCode   string
	LengthOfItems        int
	EventItem            []EventItem
	TextLength           int
	TextChar             string
}

type EventItem struct {
	ItemDescriptionLength int
	ItemDescriptionChar   string
	ItemLength            int
	ItemChar              []byte //string
}

func parseExtendedEventDescriptor(s Section) ExtendedEventDescriptor {
	lengthOfItems := s.uimsbf(48, 8).toNumber()
	textLength := s.uimsbf(56+lengthOfItems*8, 8).toNumber()
	return ExtendedEventDescriptor{
		DescriptorTag:        s.uimsbf(0, 8).toByte(),
		DescriptorLength:     s.uimsbf(8, 8).toNumber(),
		DescriptorNumber:     s.uimsbf(16, 4).toNumber(),
		LastDescriptorNumber: s.uimsbf(20, 4).toNumber(),
		ISO639LanguageCode:   s.bslbf(24, 24).toLanguageCode(),
		LengthOfItems:        lengthOfItems,
		EventItem:            parseEventItem(s[7 : 7+lengthOfItems]),
		TextLength:           textLength,
		TextChar:             s.uimsbf(64+lengthOfItems*8, textLength*8).ToString(),
	}
}

func parseEventItem(s Section) []EventItem {
	var tmp []EventItem
	for len(s) > 0 {
		itemDescriptionLength := s.uimsbf(0, 8).toNumber()
		itemLength := s.uimsbf(8+itemDescriptionLength*8, 8).toNumber()
		tmp = append(tmp, EventItem{
			ItemDescriptionLength: itemDescriptionLength,
			ItemDescriptionChar:   s.uimsbf(8, itemDescriptionLength*8).ToString(),
			ItemLength:            itemLength,
			ItemChar:              s.uimsbf(16+itemDescriptionLength*8, itemLength*8),
		})
		s = s[itemDescriptionLength+itemLength+2:]
	}
	return tmp
}

/*
Time Shifted Event Descriptor
*/

type ComponentDescriptor struct {
	DescriptorTag      byte
	DescriptorLength   int
	StreamContent      byte
	ComponentType      byte
	ComponentTag       byte
	ISO639LanguageCode string
	TextChar           string
}

func parseComponentDescriptor(s Section) ComponentDescriptor {
	charLength := len(s) - 8
	return ComponentDescriptor{
		DescriptorTag:      s.uimsbf(0, 8).toByte(),
		DescriptorLength:   s.uimsbf(8, 8).toNumber(),
		StreamContent:      s.uimsbf(20, 4).toByte(),
		ComponentType:      s.uimsbf(24, 8).toByte(),
		ComponentTag:       s.uimsbf(32, 8).toByte(),
		ISO639LanguageCode: s.bslbf(40, 24).toLanguageCode(),
		TextChar:           s.uimsbf(64, charLength*8).ToString(),
	}
}

func (d *ComponentDescriptor) GetString() string {
	return COMPONENT_TYPE[d.StreamContent][d.ComponentType]
}

/*
Mosaic Descriptor
Stream Identifier Descriptor
CA Identifier Descriptor
*/

type ContentDescriptor struct {
	DescriptorTag    byte
	DescriptorLength int
	ContentNibble    []ContentNibble
}

type ContentNibble struct {
	ContentNibbleLevel1 byte
	ContentNibbleLevel2 byte
	UserNibble1         byte
	UserNibble2         byte
}

func parseContentDescriptor(s Section) ContentDescriptor {
	return ContentDescriptor{
		DescriptorTag:    s.uimsbf(0, 8).toByte(),
		DescriptorLength: s.uimsbf(8, 8).toNumber(),
		ContentNibble:    parseContentNibble(s[2:]),
	}
}

func (c ContentNibble) ToString() (string, string) {
	return CONTENT_TYPE[c.ContentNibbleLevel1][0xff], CONTENT_TYPE[c.ContentNibbleLevel1][c.ContentNibbleLevel2]
}

func parseContentNibble(s Section) []ContentNibble {
	var tmp []ContentNibble
	for len(s) > 0 {
		tmp = append(tmp, ContentNibble{
			ContentNibbleLevel1: s.uimsbf(0, 4).toByte(),
			ContentNibbleLevel2: s.uimsbf(4, 4).toByte(),
			UserNibble1:         s.uimsbf(8, 4).toByte(),
			UserNibble2:         s.uimsbf(12, 4).toByte(),
		})
		s = s[2:]
	}
	return tmp
}

/*
Parental Rating Descriptor
Hierarchical Transmission Descriptor
Digital Copy Control Descriptor
Emergency Information Descriptor
Data Component Descriptor
System Management Descriptor
Local Time Offset Descriptor
*/

type AudioComponentDescriptor struct {
	DescriptorTag       byte
	DescriptorLength    int
	StreamContent       byte
	ComponentType       byte
	ComponentTag        byte
	StreamType          byte
	SimulcastGroupTag   byte
	ESMultiLingualFlag  bool
	MainComponentFlag   bool
	QualityIndicator    byte
	SamplingRate        byte
	ISO639LanguageCode  string
	ISO639LanguageCode2 string
	TextChar            string
}

func parseAudioComponentDescriptor(s Section) AudioComponentDescriptor {
	ESMultiLingualFlag := s.bslbf(56, 1).toBool()
	var isoCode2 string
	if ESMultiLingualFlag {
		isoCode2 = s.bslbf(88, 24).toLanguageCode()
	}
	charLength := len(s) - 8
	return AudioComponentDescriptor{
		DescriptorTag:       s.uimsbf(0, 8).toByte(),
		DescriptorLength:    s.uimsbf(8, 8).toNumber(),
		StreamContent:       s.uimsbf(20, 4).toByte(),
		ComponentType:       s.uimsbf(24, 8).toByte(),
		ComponentTag:        s.uimsbf(32, 8).toByte(),
		StreamType:          s.uimsbf(40, 8).toByte(),
		SimulcastGroupTag:   s.bslbf(48, 8).toByte(),
		ESMultiLingualFlag:  ESMultiLingualFlag,
		MainComponentFlag:   s.bslbf(57, 1).toBool(),
		QualityIndicator:    s.bslbf(58, 2).toByte(),
		SamplingRate:        s.uimsbf(60, 3).toByte(),
		ISO639LanguageCode:  s.bslbf(64, 24).toLanguageCode(),
		ISO639LanguageCode2: isoCode2,
		TextChar:            s.uimsbf(64, charLength*8).ToString(),
	}
}

/*
Target Region Descriptor
Hyperlink Descriptor
*/

type DataContentDescriptor struct {
	DescriptorTag      byte
	DescriptorLength   int
	DataComponentId    int
	EntryComponent     byte
	SelectorLength     int
	SelectorByte       []byte
	NumOfComponentRef  int
	ComponentRef       []byte
	ISO639LanguageCode string
	TextLength         int
	TextChar           string
}

func parseDataContentDescriptor(s Section) DataContentDescriptor {
	selectorLength := s.uimsbf(40, 8).toNumber()
	numOfComponentRef := s.uimsbf(48+selectorLength*8, 8).toNumber()
	textLength := s.uimsbf(80+selectorLength*8+numOfComponentRef*8, 8).toNumber()
	return DataContentDescriptor{
		DescriptorTag:      s.uimsbf(0, 8).toByte(),
		DescriptorLength:   s.uimsbf(8, 8).toNumber(),
		DataComponentId:    s.uimsbf(16, 16).toNumber(),
		EntryComponent:     s.uimsbf(32, 8).toByte(),
		SelectorLength:     selectorLength,
		SelectorByte:       s.uimsbf(48, selectorLength*8),
		NumOfComponentRef:  numOfComponentRef,
		ComponentRef:       s.uimsbf(56+selectorLength*8, numOfComponentRef*8),
		ISO639LanguageCode: s.bslbf(56+selectorLength*8+numOfComponentRef*8, 24).toLanguageCode(),
		TextLength:         textLength,
		TextChar:           s.uimsbf(88+selectorLength*8+numOfComponentRef*8, textLength).ToString(),
	}
}

/*
Video Decode Control Descriptor
Basic Local Event Descriptor
Reference Descriptor
Node Relation Descriptor
Short Node Information Descriptor
STC Reference Descriptor
Partial Reception Descriptor
Series Descriptor
*/

type EventGroupDescriptor struct {
	DescriptorTag    byte
	DescriptorLength int
	GroupType        byte
	EventCount       int
	// 実装待ち
}

/*
SI Parameter Descriptor
Broadcaster Name Descriptor
Component Group Descriptor
SI Prime TS Descriptor
Board Information Descriptor
LDT linkage Descriptor
Connected Transmission Descriptor
TS Information Descriptor
Extended Broadcaster Descriptor
Logo Transmission Descriptor
Content Availability Descriptor
Carousel Compatible Composite Descriptor
Conditional Playback Descriptor
AVC Video Descriptor
AVC timing and HRD Descriptor
Service Group Descriptor

Partial Transport Stream Descriptor
Partial Transport Stream Time Descriptor
Download Content Descriptor
CA EMM TS Descriptor
CA Contract Information Descriptor
CA Service Descriptor
Carousel Identifier Descriptor
Association Tag Descriptor
Deferred Association tags Descriptor
*/
