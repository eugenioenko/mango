package mango

type MangoReturn struct {
	From  string
	Value MangoData
	Kind  int
}

func NewMangoReturn(from string, value MangoData) *MangoReturn {
	return &MangoReturn{Kind: MangoTypeReturn, From: from, Value: value}
}

func (data *MangoReturn) ToString() string {
	return data.Value.ToString()
}

func (data *MangoReturn) ToBoolean() bool {
	return data.Value.ToBoolean()
}

func (data *MangoReturn) ToInteger() int64 {
	return data.Value.ToInteger()
}

func (data *MangoReturn) ToFloat() float64 {
	return data.Value.ToFloat()
}

func (data *MangoReturn) GetType() int {
	return data.Kind
}

func (data *MangoReturn) GetTypeName() string {
	return "return"
}

func (data *MangoReturn) GetValue() interface{} {
	return data.Value
}

func (data *MangoReturn) Equals(other MangoData) bool {
	return other.GetType() == MangoTypeFunction && data.GetValue() == other.GetValue()
}
