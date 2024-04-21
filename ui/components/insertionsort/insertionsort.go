package insertionsort

func InsertionSort(input map[int]string) map[int]string {
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		j := i - 1
		for j >= 0 && keys[j] > keys[j+1] {
			keys[j], keys[j+1] = keys[j+1], keys[j]
			j -= 1
		}
	}

	newInput := make(map[int]string)

	for _, k := range keys {
		newInput[k] = input[k]
	}
	return newInput
}
