package easyinput

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadInt(reader *bufio.Reader) (int, error) {
	line, err := (*reader).ReadString('\n')
	if err != nil {
		return 0, err // reader error
	}
	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		return 0, err // strconv error
	}
	return n, nil
}

func ReadIntLine(reader *bufio.Reader) ([]int, error) {
	line, err := (*reader).ReadString('\n')
	if err != nil {
		return nil, err // reader error
	}
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, " ")
	numbers := make([]int, len(tokens))
	for i, elem := range tokens {
		numbers[i], err = strconv.Atoi(elem)
		if err != nil {
			return numbers, err // strconv error
		}
	}
	return numbers, nil
}
