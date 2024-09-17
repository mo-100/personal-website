package utils

func GetOneToMany[T, V any, K comparable](items []T, subitems []V, getId func(T) K) ([]T, [][]V) {
	// maps each elements id to its index
	id_map := make(map[K]int)

	items_result := []T{}
	subitems_result := [][]V{}
	for i, item := range items {
		cur_idx, found := id_map[getId(item)]
		if found {
			// if the item is in the map then find its index and add child to its index
			subitems_result[cur_idx] = append(subitems_result[cur_idx], subitems[i])
		} else {
			// if the item is not in the map then add it and append it to the results
			id_map[getId(item)] = len(items_result)
			items_result = append(items_result, items[i])
			subitems_result = append(subitems_result, []V{subitems[i]})
		}
	}
	return items_result, subitems_result
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
