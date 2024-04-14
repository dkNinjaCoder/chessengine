package pieces

type Rook struct {
	GenericPiece
}

func NewRook(color Color) *Rook {
	return &Rook{
		GenericPiece{
			fenNotation: "R",
			sanNotation: "R",
			color:       color,
		},
	}
}
