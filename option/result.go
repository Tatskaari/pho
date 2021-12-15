package Option

import "errors"

type Option[T any] struct {
	Ok *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{Ok: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (r Option[T]) IsSome() bool {
	return r.Ok != nil
}

func (r Option[T]) IsNone() bool {
	return !r.IsSome()
}

func (r Option[T]) Unwrap() T {
	return *r.Ok
}

func (r Option[T]) Use(f func(T)) error {
	if r.Ok != nil {
		f(r.Unwrap())
		return nil
	}
	return errors.New("empty")
}

func (r Option[T]) OrElse(defaultVal T) T {
	if r.IsSome() {
		return r.Unwrap()
	}
	return defaultVal
}

func Map[T any, R any](r Option[T], f func(T) R) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return Some(f(r.Unwrap()))
}

func Then[T any, R any](r Option[T], f func(T) Option[R]) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return f(r.Unwrap())
}

func Cast[T any, R any](r Option[T]) Option[R] {
	if r.IsNone() {
		return None[R]()
	}
	return Some((interface{})(*r.Ok).(R))
}
