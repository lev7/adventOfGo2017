package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func getInput() string {
	input, e := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "src/adventOfGo2017/day12/test_input.txt"))
	if e != nil {
		panic(e)
	}
	return string(input)
}

func TestP1(t *testing.T) {
	input := getInput()
	expected := 6
	if p1(input) != expected {
		t.Errorf("p1(input) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	input := getInput()
	expected := 2
	if p2(input) != expected {
		t.Errorf("p2(input) != %v", expected)
	}
}
