// Package result contains a type definition for Result[T] which represents a result value. That is the value can be
// some value T, or it can be an error.
//
// Result should not be used to handle optional values. See github.com/tatskaari/pho/optoin for more information on that
package result

// Result represents a result that can either be okay, or an error
type Result[T any] struct {
	Ok  T
	Err error
}

// IsOk returns whether the result was a success
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

// IsErr returns whether the result was an error
func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

// Unwrap unwraps the result into an error tuple
func (r Result[T]) Unwrap() (T, error) {
	return r.Ok, r.Err
}

// MustUnwrap is like Unwrap but panics if the result was not a success
func (r Result[T]) MustUnwrap() T {
	if r.Err == nil {
		return r.Ok
	}
	panic(r.Err)
}

// Use applies f() to the result, returning the error if the result was not a success
func (r Result[T]) Use(f func(T)) error {
	if r.Err != nil {
		return r.Err
	}
	f(r.Ok)
	return nil
}

// MapErr maps the function onto the error type. This can be useful to attach extra context to the error message.
func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.Err == nil {
		return r
	}
	r.Err = f(r.Err)
	return r
}

// OrElse returns the value from the result, or in the error case, it returns ok
func (r Result[T]) OrElse(ok T) T {
	if r.Err == nil {
		return r.Ok
	}
	return ok
}

// Ok constructs a new ok result
func Ok[T any](ok T) Result[T] {
	return Result[T]{
		Ok: ok,
	}
}

// Err constructs a new error result.
func Err[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

// Map maps a function onto the result type. That is it applies the function f() to the result, if the result was a
// success. See Then and ThenWrap if your function needs to return a result or error tuple respectively.
func Map[T any, R any](r Result[T], f func(T) R) Result[R] {
	if r.Err == nil {
		return Ok(f(r.Ok))
	}
	return Result[R]{Err: r.Err}
}

// ThenWrap is a mapping function. It works the same as Then except f() returns an error tuple which is wrapped up into
// a result type for you
func ThenWrap[T any, R any](r Result[T], f func(T) (R, error)) Result[R] {
	if r.Err == nil {
		return Wrap(f(r.Ok))
	}
	return Result[R]{Err: r.Err}
}

// Then is a mapping function similar to Map except that it flattens the result type returned from f(), i.e.
// it returns a Result[R] instead of Result[Result[R]] that Map would have returned.
func Then[T any, R any](r Result[T], f func(T) Result[R]) Result[R] {
	if r.Err == nil {
		return f(r.Ok)
	}
	return Result[R]{Err: r.Err}
}

// Wrap wraps an error tuple into a result type
func Wrap[T any](ok T, err error) Result[T] {
	return Result[T]{Ok: ok, Err: err}
}

// Cast will attempt an interface cast on the value in a result type to convert one result type to another. This is
// useful because Go generics aren't co-variant, that is you can't pass a Result[*os.File] as a Result[io.Reader], even
// though you could pass an *os.File as an io.Reader.
func Cast[T any, R any](r Result[T]) Result[R] {
	if r.Err != nil {
		return Err[R](r.Err)
	}
	return Ok((interface{})(r.Ok).(R))
}
