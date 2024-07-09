package result

import (
	"errors"
	. "github.com/moznion/go-optional"
)

// InvalidResultErr is an error that indicates an invalid result type.
// It is used to enforce that Result objects must be created using NewOk or NewErr, not directly using Result{}.
var InvalidResultErr = errors.New("invalid result type")

// Result is a type that can hold either a success value (ok) or an error value (err).
type Result[T any, E any] struct {
	ok    Option[T]
	err   Option[E]
	valid bool
}

// NewOk creates a new Result with a success value.
func NewOk[T any, E any](value T) Result[T, E] {
	return Result[T, E]{
		ok:    Some(value),
		err:   None[E](),
		valid: true,
	}
}

// NewErr creates a new Result with an error value.
func NewErr[T any, E any](err E) Result[T, E] {
	return Result[T, E]{
		ok:    None[T](),
		err:   Some[E](err),
		valid: true,
	}
}

// MapOk creates a new Result with a new ok value, keeping the same error type.
func MapOk[T any, E any, U any](r Result[T, E], ok U) Result[U, E] {
	return MapOkAnd(r, func(t T) U {
		return ok
	})
}

func MapOkAnd[T any, E any, U any](r Result[T, E], ok func(T) U) Result[U, E] {
	if r.IsOk() {
		return NewOk[U, E](ok(r.Unwrap()))
	} else {
		return NewErr[U, E](r.UnwrapErr())
	}
}

// MapErr creates a new Result with a new err value, keeping the same ok type.
func MapErr[T any, E any, EU any](r Result[T, E], err EU) Result[T, EU] {
	if r.IsOk() {
		return NewOk[T, EU](r.Unwrap())
	} else {
		return NewErr[T, EU](err)
	}
}

func MapErrAnd[T any, E any, EU any](r Result[T, E], err func(E) EU) Result[T, EU] {
	if r.IsOk() {
		return NewOk[T, EU](r.Unwrap())
	} else {
		return NewErr[T, EU](err(r.UnwrapErr()))
	}
}

// Clone creates a new Result that is a copy of the original.
func (r Result[T, E]) Clone() Result[T, E] {
	if r.IsOk() {
		return NewOk[T, E](r.ok.Unwrap())
	} else if r.IsErr() {
		return NewErr[T, E](r.err.Unwrap())
	} else {
		panic(InvalidResultErr)
	}
}

// Unwrap returns the success value, panicking if it is not present.
func (r Result[T, E]) Unwrap() T {
	return r.ok.Unwrap()
}

// UnwrapErr returns the error value, panicking if it is not present.
func (r Result[T, E]) UnwrapErr() E {
	return r.err.Unwrap()
}

// UnwrapOrDefault returns the success value or the default value for the type if not present.
func (r Result[T, E]) UnwrapOrDefault() T {
	var def T
	return r.UnwrapOr(def)
}

// UnwrapOr returns the success value or the provided default value if not present.
func (r Result[T, E]) UnwrapOr(def T) T {
	if r.IsOk() {
		return r.ok.Unwrap()
	}
	return def
}

// UnwrapErrOr returns the error value or the provided default value if not present.
func (r Result[T, E]) UnwrapErrOr(def E) E {
	if r.IsErr() {
		return r.err.Unwrap()
	}
	return def
}

// Inspect calls the provided function with the success value if it is present.
func (r Result[T, E]) Inspect(f func(T)) Result[T, E] {
	if r.IsOk() {
		var ret = r.Unwrap()
		f(ret)
	}
	return r
}

// InspectErr calls the provided function with the error value if it is present.
func (r Result[T, E]) InspectErr(f func(E)) Result[T, E] {
	if r.IsErr() {
		var ret = r.UnwrapErr()
		f(ret)
	}
	return r
}

// Ok returns the success value as an Option.
func (r Result[T, E]) Ok() Option[T] {
	return r.ok
}

// Err returns the error value as an Option.
func (r Result[T, E]) Err() Option[E] {
	return r.err
}

// IsOk returns true if the Result contains a success value.
func (r Result[T, E]) IsOk() bool {
	if !r.valid {
		panic(InvalidResultErr)
	}
	return r.ok.IsSome()
}

// IsOkAnd returns true if the Result contains a success value and the provided function returns true for that value.
func (r Result[T, E]) IsOkAnd(f func(T) bool) bool {
	if r.IsOk() {
		return f(r.ok.Unwrap())
	}
	return false
}

// IsErr returns true if the Result contains an error value.
func (r Result[T, E]) IsErr() bool {
	if !r.valid {
		panic(InvalidResultErr)
	}
	return r.err.IsSome()
}

// IsErrAnd returns true if the Result contains an error value and the provided function returns true for that value.
func (r Result[T, E]) IsErrAnd(f func(E) bool) bool {
	if r.IsErr() {
		return f(r.UnwrapErr())
	}
	return false
}
