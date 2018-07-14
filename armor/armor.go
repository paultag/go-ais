package armor

func Encode(data []byte) (string, error) {
	ret := []rune{}
	for _, el := range data {
		var data byte = el
		if el > 32 {
			data = data + 8
		}
		data = data + 48
		ret = append(ret, rune(data))
	}
	return string(ret), nil
}

func Decode(data string) ([]byte, error) {
	ret := []byte{}
	for _, el := range data {
		data := byte(el - 48)
		if data > 40 {
			data = data - 8
		}
		ret = append(ret, data)
	}
	return ret, nil
}
