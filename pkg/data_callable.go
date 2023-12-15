package mango

type Callable func(*Interpreter, []Expression) MangoData

type MangoCallable struct {
	Function Callable
	Name     string
	Kind     int
}

func NewMangoCallable(name string, function Callable) *MangoCallable {
	return &MangoCallable{Kind: MangoTypeCallable, Name: name, Function: function}
}

func (data *MangoCallable) ToString() string {
	return data.Name
}

func (data *MangoCallable) ToBoolean() bool {
	return true
}

func (data *MangoCallable) ToInteger() int64 {
	return 0
}

func (data *MangoCallable) ToFloat() float64 {
	return 0
}

func (data *MangoCallable) GetType() int {
	return data.Kind
}

func (data *MangoCallable) GetTypeName() string {
	return "function"
}

func (data *MangoCallable) GetValue() interface{} {
	return data.Function
}

func (data *MangoCallable) Equals(other MangoData) bool {
	return other.GetType() == MangoTypeFunction && data.GetValue() == other.GetValue()
}
