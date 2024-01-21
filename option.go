package opt

type Option[T any] struct {
	present bool
	value   *T
}

func NewOption[T any](value *T) Option[T] {
	var present bool

	if value == nil {
		present = false
	} else {
		present = true
	}

	return Option[T]{present, value}
}

func Some[T any](value T) Option[T] {
	return Option[T]{present: true, value: &value}
}

func Nil[T any]() Option[T] {
	return Option[T]{present: false, value: nil}
}

func (o Option[T]) IsSome() bool {
	return o.present
}

func (o Option[T]) IsSomeAnd(f func(T) bool) bool {
	if !o.present {
		return false
	}
	return f(*o.value)
}

func (o Option[T]) IsNil() bool {
	return !o.present
}

func (o Option[T]) AsPtr() *T {
	return o.value
}

func (o Option[T]) Expect(msg string) T {
	if o.present {
		return *o.value
	} else {
		panic(msg)
	}
}

func (o Option[T]) Unwrap() T {
	return o.Expect("called `Option.Unwrap()` on a `nil` value")
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.present {
		return *o.value
	} else {
		return def
	}
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.present {
		return *o.value
	} else {
		return f()
	}
}

func (o Option[T]) UnwrapUnchecked() T {
	return *o.value
}

func (o Option[T]) IfSome(f func(T)) {
	if o.present {
		f(*o.value)
	}
}

func (o Option[T]) OkOr(err error) (*T, error) {
	if o.present {
		return o.value, nil
	} else {
		return nil, err
	}
}

func (o Option[T]) OkOrElse(err func() error) (*T, error) {
	if o.present {
		return o.value, nil
	} else {
		return nil, err()
	}
}

func (o Option[T]) Filter(p func(T) bool) Option[T] {
	if !o.present {
		return o
	}
	if p(*o.value) {
		return o
	}
	return Nil[T]()
}

func Map[T, O any](o Option[T], f func(T) O) Option[O] {
	if o.present {
		return Some(f(*o.value))
	} else {
		return Nil[O]()
	}
}

func MapOr[T, O any](o Option[T], def O, f func(T) O) O {
	if o.present {
		return f(*o.value)
	} else {
		return def
	}
}

func MapOrElse[T, O any](o Option[T], def func() O, f func(T) O) O {
	if o.present {
		return f(*o.value)
	} else {
		return def()
	}
}
