package util

import (
	"bufio"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Sum calculate the total value of an integer array
func Sum(in []int64) int64 {
	var rs int64

	rs = 0
	for i := 0; i < len(in); i++ {
		rs += in[i]
	}

	return rs
}

// CalFrequency Calculate Freq of each number on the slice.
func CalFrequency(in []int64) []float64 {
	rs := []float64{}

	total := Sum(in)
	sorted := SortInt64s(in)

	for i := 0; i < len(sorted); i++ {
		x := float64(sorted[i]) / float64(total)
		rounded := math.Round(x*100) / 100

		rs = append(rs, rounded)
	}

	return rs
}

// SortInt64s sort an array
func SortInt64s(in []int64) []int64 {
	intValues := make([]int, len(in))
	rs := make([]int64, len(in))

	for i, v := range in {
		intValues[i] = int(v)
	}

	sort.Ints(intValues)

	for i, v := range intValues {
		rs[i] = int64(v)
	}

	return rs
}

// ReadData read file data
func ReadData(filePath string) ([]int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)
	content := ""
	for {
		rl, err := readLine(r)
		if err != nil {
			break
		}
		if rl != "" {
			content += (rl + " ")
		}

	}

	arr := strings.Split(content, " ")
	arrLen := len(arr)
	var strData []string
	strData = arr[1 : arrLen-1]

	rs, err := sliceAtoi(strData)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return "", err
	}

	return strings.TrimRight(string(str), "\r\n"), nil
}

func sliceAtoi(sa []string) ([]int64, error) {
	si := make([]int64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}
		si = append(si, int64(i))
	}
	return si, nil
}
