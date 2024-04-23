package ergotool

import (
	"errors"
	"strings"

	"github.com/mlilley/go-sexpr"
)

type Footprint struct {
	s   *sexpr.Sexpr
	ref *Ref
	at  *At
}

func NewFootprintFromSexpr(s *sexpr.Sexpr) (*Footprint, error) {
	if !strings.EqualFold(s.Name(), "footprint") && !strings.EqualFold(s.Name(), "module") {
		return nil, errors.New("element not a footprint")
	}

	at, err := NewAtFromFootprint(s)
	if err != nil {
		return nil, err
	}
	ref, err := NewRefFromFootprint(s)
	if err != nil {
		return nil, err
	}

	return &Footprint{s: s, ref: ref, at: at}, nil
}

func (f *Footprint) Ref() *Ref {
	return f.ref
}

func (f *Footprint) At() *At {
	return f.at
}

func (f *Footprint) SetAt(at *At) {
	f.at = at
}
