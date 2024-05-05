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
	if len(s.Params()) != 2 && len(s.Params()) != 3 {
		return nil, errors.New("element 'at' requires 2 or 3 params")
	}

	f1, err := s.Params()[0].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 1 to be a float")
	}

	f2, err := s.Params()[1].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 2 to be a float")
	}

	if len(s.Params()) == 2 {
		str := sexpr.NewSexprStringQuoted("0", false)
		sp, err := sexpr.NewSexprParam(str)
		if err != nil {
			return nil, err
		}
		s.AddParam(2, sp)

		return &At{s: s, x: f1, y: f2, r: 0}, nil
	}

	f3, err := s.Params()[2].AsFloat()
	if err != nil {
		return nil, errors.New("element 'at' requires param 3 to be a float")
	}

	return &At{s: s, x: f1, y: f2, r: f3}, nil
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
	ss.SetValue(strconv.FormatFloat(x, 'f', -1, 64))
}

func (a *At) SetY(y float64) {
	a.y = y
	sp := a.s.Params()[1]
	ss := sp.Value().(*sexpr.SexprString)
	ss.SetValue(strconv.FormatFloat(y, 'f', -1, 64))
}

func (a *At) SetR(r float64) {
	a.r = r
	sp := a.s.Params()[2]
	ss := sp.Value().(*sexpr.SexprString)
	ss.SetValue(strconv.FormatFloat(r, 'f', -1, 64))
}

func (a *At) UpdateLocation(x, y, r float64) {
	a.SetX(x)
	a.SetY(y)
	a.SetR(r)
}

func (a *At) String() string {
	return a.s.String()
}
