package bubblesort

func BubbleSort(values []int) {
	flag := true

	size := len(values)

	for i := 0; i < size-1; i++ {
		flag = true

		for j := 0; j < size-i-1; j++ {
			if values[j] >= values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}

			if flag == true {
				break
			}
		}
	} //enf of i

}
