package pieces

import "strings"

type Color int

const (
	White Color = iota
	Black
)

type GenericPiece struct {
	fenNotation string
	sanNotation string
	color       Color
}

func (p *GenericPiece) Fen_abbrev() string {
	if p.color == White {
		return strings.ToUpper(p.fenNotation)
	} else if p.color == Black {
		return strings.ToLower(p.fenNotation)
	}
	return p.fenNotation
}

func (p *GenericPiece) San_abbrev() string {
	return p.sanNotation
}

func (p *GenericPiece) Color() Color {
	return p.color
}
