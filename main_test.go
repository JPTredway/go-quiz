package main

import (
	"testing"
)

func TestGetFile(t *testing.T) {
	{
		_, err := getFile("problems.csv")
		if err != nil {
			t.Errorf("Expected file to be read successfully, instead got err: %s", err)
		}
	}

	{
		_, err := getFile("no-file.csv")
		if err == nil {
			t.Errorf("Expected error to be thrown, instead got err == nil")
		}
	}
}

func TestGetLines(t *testing.T) {
	file, _ := getFile("problems.csv")
	_, err := getLines(file)
	if err != nil {
		t.Errorf("Expected 2D slice of strings, instead got err: %s", err)
	}
}
