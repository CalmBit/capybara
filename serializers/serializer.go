package serializers

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	SerializerField = iota
	SerializerMethod
)

type serializerField struct {
	Name     string
	Type     int
	JSONKey  string
	AsString bool
}

type Serializer struct {
	fields    []serializerField
	validator ValidatorPredicate
}

type ValidatorPredicate func(interface{}) error

// Stub function - other logic at some point
func ConstructSerializer(validator ValidatorPredicate) Serializer {
	return Serializer{fields: []serializerField{}, validator: validator}
}

func (s Serializer) AddField(name string) Serializer {
	s.fields = append(s.fields, serializerField{name, SerializerField, "", false})
	return s
}
func (s Serializer) AddFieldAsString(name string) Serializer {
	s.fields = append(s.fields, serializerField{name, SerializerField, "", true})
	return s
}

func (s Serializer) AddMethod(name string, jsonKey string) Serializer {
	s.fields = append(s.fields, serializerField{name, SerializerMethod, jsonKey, false})
	return s
}

func stringifyValue(value reflect.Value, asString bool) string {
	// Check our value's type, and serialize it if we can.
	// If it's not in this list, produce its type name so we
	// can debug it later.
	switch value.Type().Name() {
	case "string":
		return fmt.Sprintf("\"%s\"", strings.Replace(value.String(), "\"", "\\\"", -1))
	case "int":
		fallthrough
	case "int64":
		if asString {
			return fmt.Sprintf("\"%d\"", value.Int())
		} else {
			return fmt.Sprintf("%d", value.Int())
		}
	case "bool":
		if asString {
			return fmt.Sprintf("\"%t\"", value.Bool())
		} else {
			return fmt.Sprintf("%t", value.Bool())
		}
	case "Time":
		return fmt.Sprintf("\"%s\"", stringifyTime(value))
	default:
		return fmt.Sprintf("\"%s\"", value.String())
	}

}

func stringifyTime(v reflect.Value) string {
	if v.Type().Name() != "Time" {
		return ""
	}

	// Time is a little weird - since it's a struct, we have to do this nasty, nasty extra method
	// call on our value. On the bright side, it's fairly easy to set up and execute.
	return v.MethodByName("Format").Call([]reflect.Value{reflect.ValueOf(time.RFC3339)})[0].String()
}

func (s Serializer) SerializeToJSON(v interface{}) ([]byte, error) {
	err := s.validator(v)
	if err != nil {
		return nil, fmt.Errorf("serialization validation failed - %s", err.Error())
	}

	var buffer bytes.Buffer
	interfaceType := reflect.TypeOf(v)
	interfaceValue := reflect.ValueOf(v)

	buffer.WriteByte('{')
	for i, f := range s.fields {
		if f.Type == SerializerField {
			// Retrieve our field from the interface, and validate that we could find it.
			field, ok := interfaceType.FieldByName(f.Name)
			if !ok {
				return nil, fmt.Errorf("serialization failed on value %+v - field %s doesn't exist", v, f.Name)
			}

			// Write out our key to the buffer.
			buffer.WriteString(fmt.Sprintf("\"%s\":", field.Tag.Get("json")))

			// Retrieve the field's value from our current interface.
			value := interfaceValue.FieldByName(f.Name)

			buffer.WriteString(stringifyValue(value, f.AsString))

		} else if f.Type == SerializerMethod {
			// If the JSONKey is empty, we can't serialize the method - a method requires an alternative key, as
			// it has no JSON tagging.
			if f.JSONKey == "" {
				return nil, fmt.Errorf("serialization failed on value %+v - method %s has no JSONKey", v, f.Name)
			}

			// Retrieve the method from the interface, and validate that we could find it.
			method, ok := interfaceType.MethodByName(f.Name)
			if !ok {
				return nil, fmt.Errorf("serialization failed on value %+v - method %s doesn't exist", v, f.Name)
			}

			// Write out our key to the buffer.
			buffer.WriteString(fmt.Sprintf("\"%s\":", f.JSONKey))

			// Call our method (must not have any parameters).
			value := method.Func.Call([]reflect.Value{interfaceValue})[0]

			buffer.WriteString(stringifyValue(value, f.AsString))

		}
		// If we haven't hit the last field yet, add a comma.
		if i < len(s.fields)-1 {
			buffer.WriteString(", ")
		}
	}
	// End our JSON.
	buffer.WriteByte('}')

	return buffer.Bytes(), nil
}

func DeserializeFromJSON(buf []byte, v interface{}) {

}

func IsValidContainer(v interface{}) error {
	return nil
}
