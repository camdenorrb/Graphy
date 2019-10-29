package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadTensorFromText(test *testing.T) {

	tensor, err := ReadTensorFromText("[0, 1]", 2)

	if err != nil {
		test.Fatal(err)
	}

	assert.Equal(test, fmt.Sprintf("%v", tensor), "[0  1]")
}
