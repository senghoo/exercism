package allergies

func Allergies(code int) (res []string) {
	if code&1 != 0 {
		res = append(res, "eggs")
	}
	if code&2 != 0 {
		res = append(res, "peanuts")
	}
	if code&4 != 0 {
		res = append(res, "shellfish")
	}
	if code&8 != 0 {
		res = append(res, "strawberries")
	}
	if code&16 != 0 {
		res = append(res, "tomatoes")
	}
	if code&32 != 0 {
		res = append(res, "chocolate")
	}
	if code&64 != 0 {
		res = append(res, "pollen")
	}
	if code&128 != 0 {
		res = append(res, "cats")
	}
	return
}

func AllergicTo(code int, allergie string) bool {
	for _, a := range Allergies(code) {
		if a == allergie {
			return true
		}
	}
	return false
}
