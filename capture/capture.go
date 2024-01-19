package capture

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func UserInput() *Input {
	return &Input{}
}

type Input struct {
	Reader *bufio.Reader
}

func NewInput() *Input {
	return &Input{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (i *Input) ReadString() string {
	input, _ := i.Reader.ReadString('\n')
	return input
}

func (i *Input) ReadInt() (int, error) {
	input, err := i.Reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	in, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	return in, nil
}
