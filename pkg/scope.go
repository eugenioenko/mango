package mango

// Scope holds the structure of the environment of the runtime data state
type Scope struct {
	Values map[string]MangoData
	Parent *Scope
}

func NewScope(parent *Scope) *Scope {
	values := make(map[string]MangoData)
	scope := &Scope{Parent: parent, Values: values}
	return scope
}

func (scope *Scope) Get(key string) (MangoData, bool) {
	value, ok := scope.Values[key]
	if ok {
		return value, true
	}
	if scope.Parent != nil {
		return scope.Parent.Get(key)
	}
	return nil, false
}

func (scope *Scope) Has(key string) bool {
	_, ok := scope.Values[key]
	return ok
}

func (scope *Scope) Set(key string, value MangoData) {
	scope.Values[key] = value
}
