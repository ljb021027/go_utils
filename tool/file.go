package tool

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadCsv(filePath, s string) ([][]string, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	defer fi.Close()

	result := make([][]string, 0)
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		split := strings.Split(string(a), s)
		result = append(result, split)
	}
	return result, nil
}
