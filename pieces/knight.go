package pieces

type Knight struct {
	GenericPiece
}

func NewKnight(color Color) *Knight {
	return &Knight{
		GenericPiece{
			fenNotation: "N",
			sanNotation: "N",
			color:       color,
		},
	}
}
