package mango

type MangoBoolean struct {
	Value bool
	Kind  int
}

func NewMangoBoolean(value bool) *MangoBoolean {
	return &MangoBoolean{Value: value, Kind: MangoTypeBoolean}
}
func (data *MangoBoolean) ToString() string {
	if data.Value {
		return "true"
	}
	return "false"
}

func (data *MangoBoolean) ToBoolean() bool {
	return data.Value
}

func (data *MangoBoolean) ToInteger() int64 {
	// TODO: confirm if its better to return error here instead of a value
	if data.Value {
		return 1
	}
	return 0
}

func (data *MangoBoolean) ToFloat() float64 {
	if data.Value {
		return 1
	}
	return 0
}

func (data *MangoBoolean) GetType() int {
	return data.Kind
}

func (data *MangoBoolean) GetTypeName() string {
	return "boolean"
}

func (data *MangoBoolean) GetValue() interface{} {
	return data.Value
}

func (data *MangoBoolean) Equals(other MangoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
