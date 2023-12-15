package mango

import "strconv"

type MangoFloat struct {
	Value float64
	Kind  int
}

func NewMangoFloat(value float64) *MangoFloat {
	return &MangoFloat{Value: value, Kind: MangoTypeFloat}
}

func (data *MangoFloat) ToString() string {
	return strconv.FormatFloat(data.Value, 'E', -1, 64)
}

func (data *MangoFloat) ToBoolean() bool {
	return data.Value != 0
}

func (data *MangoFloat) ToInteger() int64 {
	return int64(data.Value)
}

func (data *MangoFloat) ToFloat() float64 {
	return data.Value
}

func (data *MangoFloat) GetType() int {
	return data.Kind
}

func (data *MangoFloat) GetTypeName() string {
	return "float"
}

func (data *MangoFloat) GetValue() interface{} {
	return data.Value
}

func (data *MangoFloat) Equals(other MangoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
