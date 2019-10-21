package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 12 {
		return 0
	} else if nb == 0 {
		return 1
	}

	b := 1

	for i := 1; i <= nb; nb++ {

		b = b * i
	}

	return b

}
