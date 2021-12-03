package internal

import (
	"embed"
	"fmt"
)

//go:embed inputs/*.txt
var days embed.FS

func Input(day int) (string, error) {
	b, err := days.ReadFile(fmt.Sprintf("inputs/%v.txt", day))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
