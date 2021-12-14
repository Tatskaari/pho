package result


type Result[T any] struct {
	Ok T
	Err error
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) Unwrap() (T, error) {
	return r.Ok, r.Err
}

func (r Result[T]) MustUnwrap() T {
	if r.Err == nil {
		return r.Ok
	}
	panic(r.Err)
}

func (r Result[T]) Use(f func(T)) error {
	if r.Err != nil {
		return r.Err
	}
	f(r.Ok)
	return nil
}

func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.Err == nil {
		return r
	}
	r.Err = f(r.Err)
	return r
}

func (r Result[T]) OrElse(ok T) T {
	if r.Err == nil {
		return r.Ok
	}
	return ok
}


func Ok[T any](ok T) Result[T] {
	return Result[T]{
		Ok: ok,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func Map[T any, R any](r Result[T], f func(T) R) Result[R] {
	if r.Err == nil {
		return Ok(f(r.Ok))
	}
	return Result[R]{Err: r.Err}
}

func ThenWrap[T any, R any](r Result[T], f func(T) (R, error)) Result[R] {
	if r.Err == nil {
		return Wrap(f(r.Ok))
	}
	return Result[R]{Err: r.Err}
}

func Then[T any, R any](r Result[T], f func(T) Result[R]) Result[R] {
	if r.Err == nil {
		return f(r.Ok)
	}
	return Result[R]{Err: r.Err}
}

func Wrap[T any](ok T, err error) Result[T] {
	return Result[T]{Ok: ok, Err: err}
}

func Cast[T any, R any](r Result[T]) Result[R] {
	if r.Err != nil {
		return Err[R](r.Err)
	}
	return Ok((interface{})(r.Ok).(R))
}
