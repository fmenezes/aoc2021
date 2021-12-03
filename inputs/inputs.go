package inputs

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var days embed.FS

func Input(day int) (string, error) {
	b, err := days.ReadFile(fmt.Sprintf("%v.txt", day))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
