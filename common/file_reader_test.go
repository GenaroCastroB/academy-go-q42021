package common

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("fake, csv, data")
	content, err := readFile(&buffer)
	if err != nil {
		t.Error("Failed to read csv data")
	}
	fmt.Print(content)
}
