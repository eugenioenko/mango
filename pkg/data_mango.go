package mango

type MangoData interface {
	ToString() string
	ToBoolean() bool
	ToInteger() int64
	ToFloat() float64
	Equals(other MangoData) bool
	GetType() int
	GetTypeName() string
	GetValue() interface{}
}

const (
	MangoTypeNull      = 0
	MangoTypeBoolean   = 1
	MangoTypeInteger   = 2
	MangoTypeFloat     = 3
	MangoTypeString    = 4
	MangoTypeCallable  = 5
	MangoTypeFunction  = 6
	MangoTypeReturn    = 7
	MangoTypeException = 8
)
