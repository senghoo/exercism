package queenattack

import "errors"

type queen struct {
	x, y byte
}

func newQueen(s string) (q *queen, ok bool) {
	if len(s) != 2 {
		return nil, false
	}

	a, b := s[0], s[1]

	if a < 'a' || a > 'h' || b < '0' || b > '8' {
		return nil, false
	}

	return &queen{a, b}, true

}

func (a queen) CanAttack(b queen) bool {
	// in same row or column
	if a.x == b.x || a.y == b.y {
		return true
	}
	// in diagonal
	if b.x+(a.y-b.y) == a.x {
		return true
	}
	if b.x-(a.y-b.y) == a.x {
		return true
	}
	return false
}

// CanQueenAttack check queen can attack each other
func CanQueenAttack(a, b string) (can bool, err error) {
	aq, ok := newQueen(a)
	if !ok {
		return false, errors.New("arg a invaild")
	}

	bq, ok := newQueen(b)
	if !ok {
		return false, errors.New("arg a invaild")
	}

	if aq.x == bq.x && aq.y == bq.y {
		return false, errors.New("same square")
	}

	return aq.CanAttack(*bq), nil

}
