// generated by gotemplate

package opt

import (
	"fmt"

	"github.com/Bogdan-D/easyjson/jlexer"
	"github.com/Bogdan-D/easyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Float32 struct {
	V       float32
	Defined bool
}

// Creates an optional type with a given value.
func OFloat32(v float32) Float32 {
	return Float32{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Float32) Get(deflt float32) float32 {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Float32) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Float32(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Float32) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Float32{}
	} else {
		v.V = l.Float32()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Float32) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Float32) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Float32) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Float32) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
