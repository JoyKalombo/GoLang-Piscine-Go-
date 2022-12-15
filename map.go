package main

// package piscine

func IsPrime(nbr int) bool {
	if nbr <= 1 {
		return false
	} else if nbr == 2 {
		return true
	} else {
		for i := 2; i < nbr; i++ {
			if nbr%i == 0 {
				return false
			}
			return true
		}
	}
}

func Map(f func(int) bool, a []int) []bool {
}
