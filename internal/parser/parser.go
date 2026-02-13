package parser

import (
	"database/sql"
	"fmt"
	rpsl "github.com/frederic-arr/rpsl-go"
	"os"
)

func ProcessFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}
func ParseObjects(raw string, db *sql.DB) ([]rpsl.Object, error) {
	objects, err := rpsl.ParseMany(raw)

	if err != nil {
		return nil, fmt.Errorf("Error parsing RPSL data: %v\n", err)
	}
	return objects, nil
}
