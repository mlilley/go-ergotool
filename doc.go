package ergotool

import (
	"bufio"

	"github.com/labstack/gommon/log"
	"github.com/mlilley/go-sexpr"
)

type Doc struct {
	root *sexpr.Sexpr
}

func NewDoc(r *bufio.Reader) (*Doc, error) {
	s, err := sexpr.Parse(r)
	if err != nil {
		return nil, err
	}

	return &Doc{root: s}, nil
}

func (d *Doc) Root() *sexpr.Sexpr {
	return d.root
}

func (d *Doc) EnumerateFootprints() (map[string]*Footprint, error) {
	var err error
	footprints := map[string]*Footprint{}

	err = d.enumerateFootprints("footprint", &footprints)
	if err != nil {
		return nil, err
	}

	err = d.enumerateFootprints("module", &footprints)
	if err != nil {
		return nil, err
	}

	return footprints, nil
}

func (d *Doc) enumerateFootprints(name string, footprints *map[string]*Footprint) error {
	sexprs := d.root.FindDirectChildrenByName(name)
	for _, s := range sexprs {
		fp, err := NewFootprintFromSexpr(s)
		if err != nil {
			log.Warnf("ignoring invalid footprint at Line %d, Col %d: %s", s.Line(), s.Col(), err.Error())
			continue
		}
		if _, has := (*footprints)[fp.Ref().Value()]; has {
			log.Warnf("ignoring footprint with duplicate ref '%s' at Line %d, Col %d", fp.Ref().Value(), s.Line(), s.Col())
			continue
		}
		(*footprints)[fp.Ref().Value()] = fp
	}
	return nil
}

func (d *Doc) String() string {
	return d.root.String()
}
