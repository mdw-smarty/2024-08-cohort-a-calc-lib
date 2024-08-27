package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/mdwhatcott/calc-lib/calc"
)

func TestTooFewArguments(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle(nil)
	if !errors.Is(err, errTooFewArgs) {
		t.Errorf("expected %v, got %v", errTooFewArgs, err)
	}
	if output.Len() > 0 {
		t.Error("expected no output")
	}
}
func TestInvalidFirstArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"NaN", "1"})
	if !errors.Is(err, errMalformedArgument) {
		t.Errorf("expected %v, got %v", errMalformedArgument, err)
	}
	if output.Len() > 0 {
		t.Error("expected no output")
	}
}
func TestInvalidSecondArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "NaN"})
	if !errors.Is(err, errMalformedArgument) {
		t.Errorf("expected %v, got %v", errMalformedArgument, err)
	}
	if output.Len() > 0 {
		t.Error("expected no output")
	}
}
func TestOutputWriterError(t *testing.T) {
	taco := errors.New("taco")
	output := &ErringWriter{err: taco}
	handler := NewHandler(calc.Addition{}, output)
	err := handler.Handle([]string{"1", "2"})
	if !errors.Is(err, errOutputWriteErr) {
		t.Errorf("expected %v, got %v", errOutputWriteErr, err)
	}
	if !errors.Is(err, taco) {
		t.Errorf("expected %v, got %v", taco, err)
	}
}
func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "2"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	actual := output.String()
	if actual != "3\n" {
		t.Errorf("expected %s, got %s", "3", actual)
	}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
