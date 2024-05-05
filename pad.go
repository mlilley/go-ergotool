package ergotool

import (
	"errors"
	"strings"

	"github.com/mlilley/go-sexpr"
)

type Pad struct {
	s  *sexpr.Sexpr
	at *At
}

func NewPadFromSexpr(s *sexpr.Sexpr) (*Pad, error) {
	if !strings.EqualFold(s.Name(), "pad") {
		return nil, errors.New("element not a pad")
	}

	atSexpr := s.FindDirectChildByName("at")
	if atSexpr == nil {
		return nil, errors.New("pad missing 'at'")
	}

	at, err := NewAtFromSexpr(atSexpr)
	if err != nil {
		return nil, err
	}

	return &Pad{s: s, at: at}, nil
}

func (p *Pad) ApplyRotation(r float64) {
	pr := p.at.R() + r
	p.at.SetR(pr)
}
