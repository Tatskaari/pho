package slice

import (
	"github.com/tatskaari/pho/option"
)

func Filter[T any](slice []T, pred func(T) bool) []T {
	ret := make([]T, 0, len(slice)/2)
	for _, t := range slice {
		if pred(t) {
			ret = append(ret, t)
		}
	}
	return ret
}

func FilterNil[T any](slice []*T) []*T {
	return Filter(slice, func(item *T) bool {
		return item != nil
	})
}

func FindFirst[T any](slice []T, pred func(T) bool) option.Option[T] {
	for _, item := range slice {
		if pred(item) {
			return option.Some(item)
		}
	}
	return option.None[T]()
}

func FindLast[T any](slice []T, pred func(T) bool) option.Option[T] {
	for i := range slice {
		item := slice[len(slice)-1-i]
		if pred(item) {
			return option.Some(item)
		}
	}
	return option.None[T]()
}

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t))

	}
	return ret
}

func Flatten[T any](slice [][]T) []T {
	ret := make([]T, 0, len(slice)*2)
	for _, item := range slice {
		ret = append(ret, item...)
	}
	return ret
}

func FlatMap[T any, R any](slice []T, mapper func(T) []R) []R {
	return Flatten(Map(slice, mapper))
}

func ForEach[T any](slice []T, looper func(T)) {
	for _, item := range slice {
		looper(item)
	}
}
