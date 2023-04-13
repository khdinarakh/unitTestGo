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
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	expected := []struct {
		n    string
		msg  string
		quit bool
	}{
		{"q", "", true},
		{"something", "Please enter a whole number!", false},
		{"7", "7 is a prime number!", false},
		{"8", "8 is not a prime number because it is divisible by 2!", false},
		{"0", "0 is not prime, by definition!", false},
		{"1", "1 is not prime, by definition!", false},
		{"-11", "Negative numbers are not prime, by definition!", false},
	}

	for _, v := range expected {
		input := strings.NewReader(v.n)
		scanner := bufio.NewScanner(input)
		s, b := checkNumbers(scanner)
		if s != v.msg || b != v.quit {
			t.Errorf("Unexpected message or boolean")
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	_, err := stdin.Write([]byte("1\nq\n"))
	if err != nil {
		t.Fatalf("error writing to buffer: %v", err)
	}
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}

func Test_Intro(t *testing.T) {
	expected := "Is it Prime?\n------------\n" +
		"Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "

	defer func(orig *os.File) {
		os.Stdout = orig

	}(os.Stdout)

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating pipe: %v", err)
	}

	os.Stdout = w

	intro()

	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("error reading from pipe: %v", err)
	}

	if string(out) != expected {
		t.Errorf("unexpected value - expected:  %q \n actual %q", expected, out)
	}

}

func Test_Prompt(t *testing.T) {
	expected_value := "-> "
	storeStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating pipe: %v", err)
	}
	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = storeStdout
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("error reading from pipe: %v", err)
	}

	if string(out) != expected_value {
		t.Errorf("unexpected output: \n expected: %q\n actual:   %q", expected_value, out)
	}
}
