package utils

import (
	"gorgonia.org/tensor"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadTensorFromFile(filePath string, dims ...int) (*tensor.Dense, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return ReadTensorFromText(string(bytes), dims...)
}

func ReadTensorFromText(input string, dims ...int) (*tensor.Dense, error) {

	data, err := readFloatArray(input)

	if err != nil {
		return nil, err
	}

	return tensor.New(tensor.WithShape(dims...), tensor.WithBacking(data)), nil
}

// Ex: [0, 0.1, 0.2]
func readFloatArray(input string) ([]float32, error) {

	newInput := strings.Replace(input, " ", "", -1)
	newInput = newInput[1:(len(newInput) - 1)]

	splitValues := strings.Split(newInput, ",")
	floatValues := make([]float32, len(splitValues))

	for i, v := range splitValues {

		value, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return nil, err
		}

		floatValues[i] = float32(value)
	}

	return floatValues, nil
}

/*
func ReadPointsFile(filePath string) []float32 {

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	return ReadPointsString(string(bytes))
}

func ReadPointsString(data string) []float32 {

	data = strings.Replace(data, " ", "", -1)

	output := make([]float32, 10)
	currentIndex := strings.IndexRune(data, '[') + 1

	for true {

		end := strings.IndexRune(data[currentIndex:], ']') + currentIndex
		values := strings.Split(data[currentIndex:end], ",")

		println(data[currentIndex:end])
		x, err := strconv.ParseFloat(values[0], 32)

		if err != nil {
			panic(err)
		}

		y, err := strconv.ParseFloat(values[1], 32)

		if err != nil {
			panic(err)
		}

		output[int(x)] = float32(y)

		currentIndex = strings.IndexRune(data[end:], '[') + end + 1

		if currentIndex + 1 >= len(data) {
			break
		}
	}

	return output
}
*/
