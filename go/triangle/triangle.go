//+build !example

package triangle

type Kind int

const (
	Equ = iota // equilateral
	Iso        // isosceles
	Sca        // scalene
	NaT        // not a triangle
)

// Organize your code for readability.

func KindFromSides(a, b, c float64) Kind {
	if a+b > c && a+c > b && b+c > a {
		if a == b || b == c || a == c {
			if a == b && b == c {
				return Equ
			} else {
				return Iso
			}
		} else {
			return Sca
		}
	} else {
		return NaT
	}
}
