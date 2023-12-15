package mango

import "strconv"

type MangoString struct {
	Value string
	Kind  int
}

func NewMangoString(value string) *MangoString {
	return &MangoString{Value: value, Kind: MangoTypeString}
}
func (data *MangoString) ToString() string {
	return data.Value
}

func (data *MangoString) ToBoolean() bool {
	return len(data.Value) > 0
}

func (data *MangoString) ToInteger() int64 {
	value, err := strconv.ParseInt(data.Value, 10, 64)
	if err != nil {
		panic("Cant convert string to int")
	}
	return value
}

func (data *MangoString) ToFloat() float64 {
	value, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		panic("Cant convert string to float")
	}
	return value
}

func (data *MangoString) GetType() int {
	return data.Kind
}

func (data *MangoString) GetTypeName() string {
	return "string"
}

func (data *MangoString) GetValue() interface{} {
	return data.Value
}

func (data *MangoString) Equals(other MangoData) bool {
	return other.GetType() == data.Kind && other.ToString() == data.Value
}
