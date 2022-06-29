package test

func Faktorial(nilai int) int {
	if nilai == 0 {
		return 1
	} else {
		return nilai * Faktorial(nilai-1)
	}
}
