package pieces

type Queen struct {
	GenericPiece
}

func NewQueen(color Color) *Queen {
	return &Queen{
		GenericPiece{
			fenNotation: "Q",
			sanNotation: "Q",
			color:       color,
		},
	}
}
