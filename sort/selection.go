package sort

func Selection(list []int) []int {
	for i := range list {
		for j := range list {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	return list
}
