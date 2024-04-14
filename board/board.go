package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dkNinjaCoder/chessengine/pieces"
)

func GetColor(fen rune) pieces.Color {
	if strings.ContainsAny(string(fen), "RNBQKP") {
		return pieces.White
	} else {
		return pieces.Black
	}
}

type Board struct {
	board    [][]Piece
	turn     pieces.Color
	castles  [4]bool // KQkq (white then black)
	ep       string  // en-passant target square
	halfmove int
	fullmove int
}

func NewBoard() *Board {
	b := make([][]Piece, 8)
	for i := range b {
		b[i] = make([]Piece, 8)
	}
	return &Board{
		board:    b,
		turn:     pieces.White,
		castles:  [4]bool{true, true, true, true},
		halfmove: 0,
		fullmove: 1,
	}
}

func ConvertToRankFile(row, col int) (string, error) {
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return "", fmt.Errorf("row or column out of range row: %d col: %d", row, col)
	}
	file := string("abcdefgh"[col])
	rank := string("87654321"[row])
	return file + rank, nil
}

// Returns row, col
func ConvertFromRankFile(rf string) (int, int, error) {
	file := string(rf[0])
	rank := string(rf[1])
	row, col := 0, 0
	switch file {
	case "a":
		col = 0
	case "b":
		col = 1
	case "c":
		col = 2
	case "d":
		col = 3
	case "e":
		col = 4
	case "f":
		col = 5
	case "g":
		col = 6
	case "h":
		col = 7
	default:
		return -1, -1, fmt.Errorf("unknown file %s", file)
	}
	switch rank {
	case "8":
		row = 0
	case "7":
		row = 1
	case "6":
		row = 2
	case "5":
		row = 3
	case "4":
		row = 4
	case "3":
		row = 5
	case "2":
		row = 6
	case "1":
		row = 7
	default:
		return -1, -1, fmt.Errorf("unknown rank %s", rank)
	}
	return row, col, nil
}

func (b *Board) StartingPosition() error {
	const starting_fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen_elements := strings.Split(starting_fen, " ")
	if len(fen_elements) != 6 {
		return fmt.Errorf("unexpected number of fen elements %d", len(fen_elements))
	}
	turn, castling, ep, hm, fm := fen_elements[1], fen_elements[2], fen_elements[3], fen_elements[4], fen_elements[5]
	var err error
	b.ep = ep
	b.halfmove, err = strconv.Atoi(hm)
	if err != nil {
		return fmt.Errorf("error reading halfmove %w", err)
	}
	b.fullmove, err = strconv.Atoi(fm)
	if err != nil {
		return fmt.Errorf("error reading fullmove %w", err)
	}
	switch turn {
	case "w":
		b.turn = pieces.White
	case "b":
		b.turn = pieces.Black
	default:
		return fmt.Errorf("unknown turn marker %s", turn)
	}

	b.castles = [4]bool{false, false, false, false}
	if castling == "-" {
		// nothing to do
	} else {
		for _, r := range castling {
			switch r {
			case 'K':
				b.castles[0] = true
			case 'Q':
				b.castles[1] = true
			case 'k':
				b.castles[2] = true
			case 'q':
				b.castles[3] = true
			}
		}
	}

	board_fen := strings.Split(fen_elements[0], "/")
	if len(board_fen) != 8 {
		return fmt.Errorf("unexpected board length %d", len(board_fen))
	}
	col := 0
	row := 0
	for _, item := range board_fen {
		for _, piece := range item {
			if num, err := strconv.Atoi(string(piece)); err == nil {
				// this was a number so let's skip that many cols
				col += num
			} else if strings.ToUpper(string(piece)) == "R" {
				b.board[row][col] = pieces.NewRook(GetColor(piece))
				col += 1
			} else if strings.ToUpper(string(piece)) == "N" {
				b.board[row][col] = pieces.NewKnight(GetColor(piece))
				col += 1
			} else if strings.ToUpper(string(piece)) == "B" {
				b.board[row][col] = pieces.NewBishop(GetColor(piece))
				col += 1
			} else if strings.ToUpper(string(piece)) == "Q" {
				b.board[row][col] = pieces.NewQueen(GetColor(piece))
				col += 1
			} else if strings.ToUpper(string(piece)) == "K" {
				b.board[row][col] = pieces.NewKing(GetColor(piece))
				col += 1
			} else if strings.ToUpper(string(piece)) == "P" {
				b.board[row][col] = pieces.NewPawn(GetColor(piece))
				col += 1
			} else {
				return fmt.Errorf("unknown piece %c", piece)
			}

			if col > 8 {
				return fmt.Errorf("column out of range %d", col)
			}
		}
		col = 0
		row += 1
		if row > 8 {
			return fmt.Errorf("row out of range %d", row)
		}
	}
	return nil
}

func (b *Board) GetFen() string {
	fen := ""
	for _, row := range b.board {
		blanks := 0
		for _, col := range row {
			if col != nil {
				if blanks != 0 {
					fen += strconv.Itoa(blanks)
					blanks = 0
				}
				fen += col.Fen_abbrev()
			} else {
				blanks++
			}
		}
		if blanks != 0 {
			fen += strconv.Itoa(blanks)
			blanks = 0
		}
		fen += "/"
	}
	fen = strings.TrimSuffix(fen, "/")
	fen += " "
	if b.turn == pieces.Black {
		fen += "b "
	} else {
		fen += "w "
	}

	if b.castles == [4]bool{false, false, false, false} {
		fen += "- "
	} else {
		if b.castles[0] {
			fen += "K"
		}
		if b.castles[1] {
			fen += "Q"
		}
		if b.castles[2] {
			fen += "k"
		}
		if b.castles[3] {
			fen += "q"
		}
		fen += " "
	}

	fen += b.ep + " " + strconv.Itoa(b.halfmove) + " " + strconv.Itoa(b.fullmove)
	return fen
}

type Piece interface {
	Fen_abbrev() string
	San_abbrev() string
	Color() pieces.Color
}
