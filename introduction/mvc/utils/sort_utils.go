package utils

// []int {3,0,7,6,5}
func BubbleSort(elements []int) []int {

	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i], elements[i+1]
				keepRunning = true
			}
		}

	}
	return elements
}
