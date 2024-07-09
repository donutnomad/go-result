package result

import (
	"testing"
)

// Unit tests
func TestResult(t *testing.T) {
	okResult := NewOk[int, string](42)
	errResult := NewErr[int, string]("error")

	// Test NewOk and NewErr
	if !okResult.IsOk() || okResult.IsErr() {
		t.Errorf("Expected okResult to be Ok")
	}
	if !errResult.IsErr() || errResult.IsOk() {
		t.Errorf("Expected errResult to be Err")
	}

	// Test Unwrap and UnwrapErr
	if okResult.Unwrap() != 42 {
		t.Errorf("Expected okResult to unwrap to 42")
	}
	if errResult.UnwrapErr() != "error" {
		t.Errorf("Expected errResult to unwrap to 'error'")
	}

	// Test UnwrapOrDefault
	defaultResult := Result[int, string]{}
	if defaultResult.UnwrapOrDefault() != 0 {
		t.Errorf("Expected defaultResult to unwrap to default int value 0")
	}

	// Test UnwrapOr
	if okResult.UnwrapOr(99) != 42 {
		t.Errorf("Expected okResult to unwrap to 42")
	}
	if defaultResult.UnwrapOr(99) != 99 {
		t.Errorf("Expected defaultResult to unwrap to 99")
	}

	// Test UnwrapErrOr
	if errResult.UnwrapErrOr("default") != "error" {
		t.Errorf("Expected errResult to unwrap to 'error'")
	}
	if defaultResult.UnwrapErrOr("default") != "default" {
		t.Errorf("Expected defaultResult to unwrap to 'default'")
	}

	// Test Inspect and InspectErr
	var inspectedValue int
	okResult.Inspect(func(val int) {
		inspectedValue = val
	})
	if inspectedValue != 42 {
		t.Errorf("Expected inspected value to be 42")
	}

	var inspectedError string
	errResult.InspectErr(func(err string) {
		inspectedError = err
	})
	if inspectedError != "error" {
		t.Errorf("Expected inspected error to be 'error'")
	}

	// Test Clone
	clonedResult := okResult.Clone()
	if !clonedResult.IsOk() || clonedResult.Unwrap() != 42 {
		t.Errorf("Expected clonedResult to be Ok with value 42")
	}

}

// Test MapOk
func TestMapOk(t *testing.T) {
	okResult := NewOk[int, string](42)
	errResult := NewErr[int, string]("error")

	newOkResult := MapOk(okResult, "new ok")
	if !newOkResult.IsOk() || newOkResult.Unwrap() != "new ok" {
		t.Errorf("Expected newOkResult to be Ok with value 'new ok'")
	}
	if newOkResult.IsErr() {
		t.Errorf("Expected newOkResult to have no error")
	}

	newErrResult := MapOk(errResult, "new ok")
	if newErrResult.IsOk() {
		t.Errorf("Expected newErrResult to have no success value")
	}
	if !newErrResult.IsErr() || newErrResult.UnwrapErr() != "error" {
		t.Errorf("Expected newErrResult to be Err with value 'error'")
	}
}

// Test MapErr
func TestMapErr(t *testing.T) {
	okResult := NewOk[int, string](42)
	errResult := NewErr[int, string]("error")

	newOkResult := MapErr(okResult, "new error")
	if !newOkResult.IsOk() || newOkResult.Unwrap() != 42 {
		t.Errorf("Expected newOkResult to be Ok with value 42")
	}
	if newOkResult.IsErr() {
		t.Errorf("Expected newOkResult to have no error")
	}

	newErrResult := MapErr(errResult, "new error")
	if newErrResult.IsOk() {
		t.Errorf("Expected newErrResult to have no success value")
	}
	if !newErrResult.IsErr() || newErrResult.UnwrapErr() != "new error" {
		t.Errorf("Expected newErrResult to be Err with value 'new error'")
	}
}
