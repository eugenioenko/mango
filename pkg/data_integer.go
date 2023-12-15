package mango

import "strconv"

type MangoInteger struct {
	Value int64
	Kind  int
}

func NewMangoInteger(value int64) *MangoInteger {
	return &MangoInteger{Value: value, Kind: MangoTypeInteger}
}

func (data *MangoInteger) ToString() string {
	return strconv.FormatInt(data.Value, 10)
}

func (data *MangoInteger) ToBoolean() bool {
	return data.Value != 0
}

func (data *MangoInteger) ToInteger() int64 {
	return data.Value
}

func (data *MangoInteger) ToFloat() float64 {
	return float64(data.Value)
}

func (data *MangoInteger) GetType() int {
	return data.Kind
}

func (data *MangoInteger) GetTypeName() string {
	return "integer"
}

func (data *MangoInteger) GetValue() interface{} {
	return data.Value
}

func (data *MangoInteger) Equals(other MangoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
