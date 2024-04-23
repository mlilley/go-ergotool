package ergotool

import (
	"errors"
	"strconv"

	"github.com/mlilley/go-sexpr"
)

type At struct {
	s *sexpr.Sexpr
	x float64
	y float64
	r float64
}

func NewAtFromAt(fromAt *At, x float64, y float64, r float64) *At {
	return &At{
		s: fromAt.s,
		x: x,
		y: y,
		r: r,
	}
}

func NewAtFromSexpr(s *sexpr.Sexpr) (*At, error) {
	if s.Name() != "at" {
		return nil, errors.New("element is not an 'at'")
	}
	if len(s.Params()) != 3 {
		return nil, errors.New("element 'at' requires 3 params")
	}

	f1, err := s.Params()[0].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 1 to be a float")
	}

	f2, err := s.Params()[1].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 2 to be a float")
	}

	f3, err := s.Params()[2].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 3 to be a float")
	}

	return &At{s: s, x: f1, y: f2, r: f3}, nil
}

func NewAtFromFootprint(s *sexpr.Sexpr) (*At, error) {
	sat := s.FindDirectChildByName("at")
	if sat == nil {
		return nil, errors.New("missing 'at'")
	}
	return NewAtFromSexpr(sat)
}

func (a *At) Sexpr() *sexpr.Sexpr {
	return a.s
}

func (a *At) X() float64 {
	return a.x
}

func (a *At) Y() float64 {
	return a.y
}

func (a *At) R() float64 {
	return a.r
}

func (a *At) SetX(x float64) {
	a.x = x
	s := a.s
	sp := s.Params()[0]
	ss := sp.Value().(*sexpr.SexprString)
	ss.Set(strconv.FormatFloat(x, 'f', -1, 64))
}

func (a *At) SetY(y float64) {
	a.y = y
	sp := a.s.Params()[1]
	ss := sp.Value().(*sexpr.SexprString)
	ss.Set(strconv.FormatFloat(y, 'f', -1, 64))
}

func (a *At) SetR(r float64) {
	a.r = r
	sp := a.s.Params()[2]
	ss := sp.Value().(*sexpr.SexprString)
	ss.Set(strconv.FormatFloat(r, 'f', -1, 64))
}

func (a *At) String() string {
	return a.s.String()
}
