package ergotool

import (
	"errors"
	"strings"

	"github.com/mlilley/go-sexpr"
)

type Ref struct {
	s *sexpr.Sexpr
	v string
}

func NewRefFromSexpr(s *sexpr.Sexpr) (*Ref, error) {
	if !strings.EqualFold(s.Name(), "property") && !strings.EqualFold(s.Name(), "fp_text") {
		return nil, errors.New("element is not a 'ref'")
	}

	if len(s.Params()) < 2 {
		return nil, errors.New("element is not a 'ref' (at least 2 params required)")
	}

	s1, err := s.Params()[0].AsString()
	if err != nil {
		return nil, errors.New("element is not a 'ref' (param 1 must be a string)")
	}
	if !strings.EqualFold(s1, "Reference") {
		return nil, errors.New("element is not a 'ref' (param 1 must have value \"Reference\")")
	}

	s2, err := s.Params()[1].AsString()
	if err != nil {
		return nil, errors.New("element is not a 'ref' (param 2 must be a string)")
	}

	return &Ref{s: s, v: s2}, nil
}

func NewRefFromFootprint(s *sexpr.Sexpr) (*Ref, error) {
	sref := s.FindChild(func(ss *sexpr.Sexpr, depth int) bool {
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

	if sref == nil {
		return nil, errors.New("missing 'ref'")
	}

	return NewRefFromSexpr(sref)
}

func (kr *Ref) Sexpr() *sexpr.Sexpr {
	return kr.s
}

func (kr *Ref) Value() string {
	return kr.v
}
