package stream

import "github.com/tatskaari/pho/option"

type metadata interface {
	Size() int
}

type Source[T any] interface {
	metadata
	Next() option.Option[T]
}

func Slice[T any](slice []T) *SliceSource[T] {
	return &SliceSource[T]{data: slice}
}

type SliceSource[T any] struct {
	data []T
	i    int
}

func (slice *SliceSource[T]) Size() int {
	return len(slice.data)
}

func (slice *SliceSource[T]) Next() option.Option[T] {
	if slice.i == len(slice.data) {
		return option.None[T]()
	}
	slice.i++
	return option.Some(slice.data[slice.i-1])
}

type mapSource[T any, R any] struct {
	metadata
	s  Source[T]
	op func(T) R
}

func (s mapSource[T, R]) Next() option.Option[R] {
	return option.Map(s.s.Next(), s.op)
}

func Map[T any, R any](source Source[T], op func(T) R) Source[R] {
	return mapSource[T, R]{
		metadata: source,
		s:        source,
		op:       op,
	}
}

type filter[T any] struct {
	s    Source[T]
	pred func(T) bool
}

func (s filter[T]) Next() option.Option[T] {
	for next := s.s.Next(); next.IsSome(); next = s.s.Next() {
		if s.pred(next.Unwrap()) {
			return next
		}
	}
	return option.None[T]()
}

func (s filter[T]) Size() int {
	return 0 // We don't know the size any more so let Go's slice resizing handle things when we collect
}

func Filter[T any](source Source[T], pred func(T) bool) Source[T] {
	return filter[T]{
		s:    source,
		pred: pred,
	}
}

func Collect[R any](s Source[R]) []R {
	ret := make([]R, 0, s.Size())
	for next := s.Next(); next.IsSome(); next = s.Next() {
		ret = append(ret, next.Unwrap())
	}
	return ret
}
