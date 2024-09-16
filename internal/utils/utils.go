package utils

func GetOneToMany[T, V any](items []T, subitems []V, get_id func(T) string) ([]T, [][]V) {
	if len(items) == 0 {
		return []T{}, [][]V{}
	}

	idx := -1
	t_list := []T{}
	v_list := [][]V{}
	for i, t := range items {
		if i == 0 || get_id(t) != get_id(t_list[idx]) {
			idx += 1
			v_list = append(v_list, []V{subitems[i]})
			t_list = append(t_list, t)
		} else {
			v_list[idx] = append(v_list[idx], subitems[i])
		}
	}
	return t_list, v_list
}

func Map[T, V any](items []T, f func(T) V) []V {
	res := make([]V, len(items))
	for i, v := range items {
		res[i] = f(v)
	}
	return res
}

func Filter[T any](items []T, f func(T) bool) []T {
	res := []T{}
	for _, t := range items {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}
