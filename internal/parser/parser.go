package parser

import (
	"fmt"
	"os"

	rpsl "github.com/frederic-arr/rpsl-go"
)

type parsedRPSL struct {
	Attribute []rpsl.Attribute
}

func ProcessFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}
func ParseObjects(raw string) (*rpsl.Object, error) {
	objects, err := rpsl.ParseMany(raw)

	if err != nil {
		return nil, fmt.Errorf("Error parsing RPSL data: %v\n", err)
	}

	for _, obj := range objects {
		return &obj, nil
	}

	return nil, fmt.Errorf("Error parsing RPSL objects: %v\n", objects)
}
