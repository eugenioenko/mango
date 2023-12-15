package mango

type MangoException struct {
	Value string
	Kind  int
}

func NewMangoException(value string) *MangoException {
	return &MangoException{Value: value, Kind: MangoTypeException}
}
func (data *MangoException) ToString() string {
	return data.Value
}

func (data *MangoException) ToBoolean() bool {
	return true
}

func (data *MangoException) ToInteger() int64 {
	return 0
}

func (data *MangoException) ToFloat() float64 {
	return 0.0
}

func (data *MangoException) GetType() int {
	return data.Kind
}

func (data *MangoException) GetTypeName() string {
	return "exception"
}

func (data *MangoException) GetValue() interface{} {
	return data.Value
}

func (data *MangoException) Equals(other MangoData) bool {
	return other.GetType() == data.Kind && other.ToString() == data.Value
}
