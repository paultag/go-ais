package sixbit

//
func Decode(data []byte) (*BitSlice, error) {
	b := NewBitSlice()
	for _, el := range data {
		data := el
		if data > 64 {
			data = data - 64
		}
		b.AppendByte(data, 6)
	}
	return &b, nil
}

//
func Encode(b *BitSlice) (string, error) {
	return "", nil
}
