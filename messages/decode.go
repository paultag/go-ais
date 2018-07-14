package messages

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"pault.ag/go/ais/sixbit"
)

// Interface to override the default unpacking method using
// struct tags.
type Unmarshallable interface {
	// Given the input bits (in the form of a sixbit.BitSlice), unpack
	// those into the struct.
	UnmarshalBits(bits *sixbit.BitSlice) error
}

// Unpack the provided BitSlice into an object.
//
// Rules:
//
func Unmarshal(bits *sixbit.BitSlice, into interface{}) error {
	return unmarshal(bits, reflect.ValueOf(into))
}

//
func unmarshal(bits *sixbit.BitSlice, into reflect.Value) error {
	if into.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("Decode can only decode into a pointer!")
	}

	switch into.Elem().Type().Kind() {
	case reflect.Struct:
		return decodeStruct(bits, into)
	default:
		return fmt.Errorf("Can't Decode into a %s", into.Elem().Type().Name())
	}

	return nil
}

//
func decodeStruct(bits *sixbit.BitSlice, into reflect.Value) error {
	/* If we have a pointer, let's follow it */
	if into.Type().Kind() == reflect.Ptr {
		return decodeStruct(bits, into.Elem())
	}

	/* If it has an Unmarshallable interface, let's go ahead and handle it. */
	if unmarshal, ok := into.Addr().Interface().(Unmarshallable); ok {
		return unmarshal.UnmarshalBits(bits)
	}

	/* Right, now, we're going to decode a Paragraph into the struct */

	for i := 0; i < into.NumField(); i++ {
		field := into.Field(i)
		fieldType := into.Type().Field(i)

		if field.Type().Kind() == reflect.Struct {
			/* If we have a nested struct, go ahead and go into it */
			err := decodeStruct(bits, field)
			if err != nil {
				return err
			}
			continue
		}

		var from uint64 = 0
		var length uint64 = 0
		var err error

		if it := fieldType.Tag.Get("bits"); it != "" {
			data := strings.SplitN(it, ":", 2)
			from, err = strconv.ParseUint(data[0], 10, 16)
			if err != nil {
				return err
			}
			length, err = strconv.ParseUint(data[1], 10, 16)
			if err != nil {
				return err
			}
		}

		slice := bits.Slice(uint(from), uint(length))

		if err := decodeStructValue(field, fieldType, slice); err != nil {
			return err
		}
	}

	return nil
}

//
func getDivisor(field reflect.Value, fieldType reflect.StructField) int64 {
	if it := fieldType.Tag.Get("divisor"); it != "" {
		val, err := strconv.ParseInt(it, 10, 64)
		if err != nil {
			panic(err)
		}
		return val
	}
	return 1
}

//
func decodeStructValue(field reflect.Value, fieldType reflect.StructField, bits *sixbit.BitSlice) error {
	switch field.Type().Kind() {
	case reflect.String:
		field.SetString(bits.StringTrim())
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		field.SetUint(bits.Uint64() / uint64(getDivisor(field, fieldType)))
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		field.SetInt(bits.Int64() / int64(getDivisor(field, fieldType)))
		return nil
	case reflect.Float32, reflect.Float64:
		field.SetFloat(float64(bits.Int64()) / float64(getDivisor(field, fieldType)))
		return nil
	case reflect.Bool:
		field.SetBool(bits.Bool())
		return nil
	default:
		return fmt.Errorf("Type not supported: %s", field.Type().Kind())
	}
}
