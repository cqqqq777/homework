package main01

func Cla(n1 float64, n2 float64, operator string) float64 {
	var res float64
	switch operator {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	}
	return res
}