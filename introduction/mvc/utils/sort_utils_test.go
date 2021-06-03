package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 8, 7, 4, 5}

	elements = BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, []int{4, 5, 7, 8, 9}, elements)

	for i := 0; i < len(elements); i++ {
		fmt.Printf("%v", elements[i])
	}

}
