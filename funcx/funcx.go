package funcx

import "github.com/ldeng7/go-x/common"

func IntrincsMax[T common.IntrincsOrd](arr ...T) T {
	return SliceMaxIntrincs(arr)
}

func IntrincsMax2[T common.IntrincsOrd](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func IntrincsMin[T common.IntrincsOrd](arr ...T) T {
	return SliceMinIntrincs(arr)
}

func IntrincsMin2[T common.IntrincsOrd](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func IntrincsLess[T common.IntrincsOrd](a, b T) bool {
	return a < b
}
