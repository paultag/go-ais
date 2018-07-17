package ais

import (
	"pault.ag/go/ais/armor"
	"pault.ag/go/ais/messages"
	"pault.ag/go/ais/sixbit"
)

//
type Message struct {
	//
	Header messages.Header

	//
	Bits *sixbit.BitSlice
}

//
func (m Message) Unmarshal(target interface{}) error {
	return messages.Unmarshal(m.Bits, target)
}

func (m Message) Parse() (interface{}, error) {
	switch m.Header.Type {
	case 1, 2, 3:
		return m.Position()
	case 5:
		return m.Voyage()
	case 18:
		return m.ClassBPosition()
	case 21:
		return m.NavigationAid()
	case 24:
		return m.StaticData()
	case 4:
		return m.BaseStation()
	default:
		/* XXX: this is bad. return an error */
		return nil, nil
	}
}

//
func (m Message) BaseStation() (*messages.BaseStation, error) {
	bs := messages.BaseStation{}
	if err := m.Unmarshal(&bs); err != nil {
		return nil, err
	}
	return &bs, nil
}

//
func (m Message) ClassBPosition() (*messages.ClassBPosition, error) {
	cbp := messages.ClassBPosition{}
	if err := m.Unmarshal(&cbp); err != nil {
		return nil, err
	}
	return &cbp, nil
}

//
func (m Message) NavigationAid() (*messages.NavigationAid, error) {
	na := messages.NavigationAid{}
	if err := m.Unmarshal(&na); err != nil {
		return nil, err
	}
	return &na, nil
}

//
func (m Message) Position() (*messages.Position, error) {
	p := messages.Position{}
	if err := m.Unmarshal(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

//
func (m Message) StaticData() (*messages.StaticData, error) {
	sd := messages.StaticData{}
	if err := m.Unmarshal(&sd); err != nil {
		return nil, err
	}
	return &sd, nil
}

//
func (m Message) Voyage() (*messages.Voyage, error) {
	voy := messages.Voyage{}
	if err := m.Unmarshal(&voy); err != nil {
		return nil, err
	}
	return &voy, nil
}

//
func (m *Message) UnmarshalBits(bits *sixbit.BitSlice) error {
	if err := messages.Unmarshal(bits, &m.Header); err != nil {
		return err
	}
	m.Bits = bits
	return nil
}

//
func Decode(data string) (*Message, error) {
	sixbytes, err := armor.Decode(data)
	if err != nil {
		return nil, err
	}

	slice, err := sixbit.Decode(sixbytes)
	if err != nil {
		return nil, err
	}

	ret := Message{}
	if err := messages.Unmarshal(slice, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
