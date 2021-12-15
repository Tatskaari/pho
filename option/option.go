// Package option contains a type definition for Option[T] which represents an optional value. That is the value can be
// some value T, or it can be empty.
//
// Option should not be used to handle errors. See github.com/tatskaari/pho/result for more information on handling
// error cases.
package option

import "errors"

type Option[T any] struct {
	Ok *T
}

// Some returns an Option of some value
func Some[T any](value T) Option[T] {
	return Option[T]{Ok: &value}
}

// None returns an empty optional
func None[T any]() Option[T] {
	return Option[T]{}
}

// IsSome returns true when the optional contains some value
func (r Option[T]) IsSome() bool {
	return r.Ok != nil
}

// IsNone returns true when the optional is empty
func (r Option[T]) IsNone() bool {
	return !r.IsSome()
}

// Unwrap returns the value from the optional. Will panic if the optional is empty. See OrElse if you want to get a
// default value for empty optionals.
func (r Option[T]) Unwrap() T {
	return *r.Ok
}

// Use calls f() on the value from the optional. Returns an error if the optional was empty.
func (r Option[T]) Use(f func(T)) error {
	if r.Ok != nil {
		f(r.Unwrap())
		return nil
	}
	return errors.New("empty")
}

// OrElse returns the value of the optional when it has one. Defaults to defaultVal when the optional is empty.
func (r Option[T]) OrElse(defaultVal T) T {
	if r.IsSome() {
		return r.Unwrap()
	}
	return defaultVal
}

// Map calls f() on the value in the optional if there is one, returning a new optional based on the return type of f().
// See Then if your function needs to return an optional itself.
func Map[T any, R any](r Option[T], f func(T) R) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return Some(f(r.Unwrap()))
}

// Then is a mapping function, like Map, that applies f to the value in the optional. Unlike map, the f() returns an
// optional, which is flattened for you. That is, Then will produce an Option[R], where Map would produce an
// Option[Option[R]]
func Then[T any, R any](r Option[T], f func(T) Option[R]) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return f(r.Unwrap())
}

// Cast will attempt an interface cast on the value in a result type to convert one result type to another. This is
// useful because Go generics aren't co-variant, that is you can't pass a Result[*os.File] as a Result[io.Reader], even
// though you could pass an *os.File as an io.Reader.
func Cast[T any, R any](r Option[T]) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return Some((interface{})(*r.Ok).(R))
}
