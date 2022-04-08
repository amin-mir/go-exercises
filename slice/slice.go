package slice

func Reverse[T any](s []T) []T {
	res := make([]T, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res[i] = s[i]
	}
	return res
}

func ReverseInPlace[S ~[]T, T any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
