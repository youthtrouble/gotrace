package trace

import "io"

//Tracer is the interface
type Tracer interface {
	Trace(...interface{})
}

//New func
func New(w io.Writer) Tracer {
	return nil
}