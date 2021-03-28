package table

type Section []byte

func (s Section) TableId() byte {
	return s[0]
}

func (s Section) SectionSyntaxIndicator() byte {
	return s[1] & 0x80 >> 7
}

func (s Section) SectionLength() int {
	return int(uint16(s[2]) | uint16(s[1]&0x0F)<<8)
}
