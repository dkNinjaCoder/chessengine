package pieces

type Bishop struct {
	GenericPiece
}

func NewBishop(color Color) *Bishop {
	return &Bishop{
		GenericPiece{
			fenNotation: "B",
			sanNotation: "B",
			color:       color,
		},
	}
}
