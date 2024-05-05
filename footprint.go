package ergotool

import (
	"errors"
	"strings"

	"github.com/mlilley/go-sexpr"
)

type Footprint struct {
	s    *sexpr.Sexpr
	ref  *Ref
	at   *At
	pads []*Pad
}

func NewFootprintFromSexpr(s *sexpr.Sexpr) (*Footprint, error) {
	if !strings.EqualFold(s.Name(), "footprint") && !strings.EqualFold(s.Name(), "module") {
		return nil, errors.New("element not a footprint")
	}

	at, err := __footprint_getAt(s)
	if err != nil {
		return nil, err // todo return error specific to this footprint
	}

	ref, err := __footprint_getRef(s)
	if err != nil {
		return nil, err // todo return error specific to this footprint
	}

	pads, err := __footprint_getPads(s)
	if err != nil {
		return nil, err
	}

	return &Footprint{s: s, ref: ref, at: at, pads: pads}, nil
}

func __footprint_getAt(s *sexpr.Sexpr) (*At, error) {
	atSexpr := s.FindDirectChildByName("at")
	if atSexpr == nil {
		return nil, errors.New("missing 'at'")
	}
	at, err := NewAtFromSexpr(atSexpr)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func __footprint_getRef(s *sexpr.Sexpr) (*Ref, error) {
	refSexpr := s.FindChild(func(ss *sexpr.Sexpr, depth int) bool {
		if !strings.EqualFold(ss.Name(), "property") && !strings.EqualFold(ss.Name(), "fp_text") {
			return false
		}
		if len(ss.Params()) < 2 {
			return false
		}
		p0, err := ss.Params()[0].AsString()
		if err != nil {
			return false
		}
		if !strings.EqualFold(p0, "reference") {
			return false
		}
		return true
	}, 1)

	if refSexpr == nil {
		return nil, errors.New("missing 'ref'")
	}

	ref, err := NewRefFromSexpr(refSexpr)
	if err != nil {
		return nil, err
	}

	return ref, nil
}

func __footprint_getPads(s *sexpr.Sexpr) ([]*Pad, error) {
	pads := []*Pad{}
	padSexprs := s.FindDirectChildrenByName("pad")
	for _, ps := range padSexprs {
		pad, err := NewPadFromSexpr(ps)
		if err != nil {
			return nil, err
		}
		pads = append(pads, pad)
	}
	return pads, nil
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

func (f *Footprint) Pads() []*Pad {
	return f.pads
}

func (f *Footprint) UpdateLocation(x, y, r float64) {
	f.at.UpdateLocation(x, y, r)
	for _, pad := range f.pads {
		pad.ApplyRotation(r)
	}
}
