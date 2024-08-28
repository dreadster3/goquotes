package lists

func Map[T, U any](arr []T, f func(T) U) []U {
	res := make([]U, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}
