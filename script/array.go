package script

import (
	"strings"

	qlova "github.com/qlova/script"
)

//Array is a dynamic-content array.
type Array struct {
	Q Ctx
	qlova.Native
}

//Strings creates a new script.Array of strings.
func (q Ctx) Strings(slice ...string) Array {
	var converted = make([]qlova.Value, len(slice))
	for i := range slice {
		converted[i] = q.String(slice[i])
	}
	return q.NewArray(converted...)
}

//NewArray creates a new Array.
func (q Ctx) NewArray(elements ...qlova.Value) Array {
	var raw = make([]string, len(elements))
	for i := range elements {
		raw[i] = q.Raw(elements[i].T())
	}
	if len(elements) > 0 {
		return Array{
			Q:      q,
			Native: q.Value("[" + strings.Join(raw, ",") + "]").Native(),
		}
	}
	return Array{
		Q:      q,
		Native: q.Value("[]").Native(),
	}
}

//Push pushes a new value to the array.
func (a Array) Push(v qlova.Type) {
	var q = a.Q
	q.Javascript(q.Raw(a.Native) + `.push(` + q.Raw(v) + ");")
}

//Index returns the value at the given index in the array.
func (a Array) Index(i Int) Dynamic {
	var q = a.Q
	return q.Value(q.Raw(a.Native) +
		`[` + q.Raw(i) + "]").Dynamic()
}

//Mutate sets the value at the given index in the array.
func (a Array) Mutate(i Int, v qlova.Type) {
	var q = a.Q
	q.Javascript(q.Raw(a.Native) +
		`[` + q.Raw(i) + "] = " +
		q.Raw(v) + ";")
}

//Var calls Native.Var(...string).
func (a Array) Var(name ...string) Array {
	var variable = a.Native
	variable.Var(name...)
	return Array{
		Q:      a.Q,
		Native: variable,
	}
}
