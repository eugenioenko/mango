package mango

type MangoFunction struct {
	Name string
	Args []string
	Body []Expression
	Kind int
}

func NewMangoFunction(name string, args []string, body []Expression) *MangoFunction {
	return &MangoFunction{Kind: MangoTypeFunction, Name: name, Args: args, Body: body}
}

func (data *MangoFunction) ToString() string {
	return data.Name
}

func (data *MangoFunction) ToBoolean() bool {
	return true
}

func (data *MangoFunction) ToInteger() int64 {
	return 0
}

func (data *MangoFunction) ToFloat() float64 {
	return 0
}

func (data *MangoFunction) GetType() int {
	return data.Kind
}

func (data *MangoFunction) GetTypeName() string {
	return "function"
}

func (data *MangoFunction) GetValue() interface{} {
	return data.Body
}

func (data *MangoFunction) Equals(other MangoData) bool {
	return other.GetType() == MangoTypeFunction && data.GetValue() == other.GetValue()
}
