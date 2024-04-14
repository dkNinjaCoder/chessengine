package pieces

type King struct {
	GenericPiece
}

func NewKing(color Color) *King {
	return &King{
		GenericPiece{
			fenNotation: "K",
			sanNotation: "K",
			color:       color,
		},
	}
}
