package table

func GetSid(section []byte) (sid int) {
	s := Section(section)
	return s.uimsbf(96, 16).toNumber()
}
