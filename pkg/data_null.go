package mango

type MangoNull struct {
	Kind int
}

func NewMangoNull() *MangoNull {
	return &MangoNull{Kind: MangoTypeNull}
}

func (data *MangoNull) ToString() string {
	return "null"
}

func (data *MangoNull) ToBoolean() bool {
	return false
}

func (data *MangoNull) ToInteger() int64 {
	return 0
}

func (data *MangoNull) ToFloat() float64 {
	return 0
}

func (data *MangoNull) GetType() int {
	return data.Kind
}

func (data *MangoNull) GetTypeName() string {
	return "null"
}

func (data *MangoNull) GetValue() interface{} {
	panic("Cant GetValue of Null")
}

func (data *MangoNull) Equals(other MangoData) bool {
	return data.GetType() == other.GetType()
}
