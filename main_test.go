package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not prime becuase it id divisible by 2"},
		{"zero", 0, false, "0 is not prime by definition"},
		{"one", 1, false, "1 is not prime by definition"},
		{"negative", -1, false, "Negative numbers are not prime"},
	}
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true, got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false, got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected message: %s , got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("prompt doesn't return correct value, expected value: %s got %s", "-> ", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Enter a number") {
		t.Errorf("intro doesn't return correct value, got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a number"},
		{"not Number", "sdq", "Please enter a number"},
		{"decimal", "1.4", "Please enter a number"},
		{"prime", "7", "7 is a prime number"},
		{"not prime", "8", "8 is not prime becuase it id divisible by 2"},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
	}
	for _, e := range tests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.Contains(res, e.expected) {
			t.Errorf("%s doesn't return correct value. expected: %s got: %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	defer close(doneChan)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan

}
