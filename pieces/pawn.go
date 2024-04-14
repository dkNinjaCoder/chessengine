package pieces

type Pawn struct {
	GenericPiece
}

func NewPawn(color Color) *Pawn {
	return &Pawn{
		GenericPiece{
			fenNotation: "P",
			sanNotation: "",
			color:       color,
		},
	}
}
