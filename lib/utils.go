package lib

func RemoveIndex[S any](slice []S, i int) []S {
	ret := make([]S, 0)
	ret = append(ret, slice[:i]...)
	return append(ret, slice[i+1:]...)
}
