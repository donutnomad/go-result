package result

func ResultScope[T any, E any]() RH[T, E] {
	return RH[T, E]{}
}

type RH[T any, E any] struct{}

func (r *RH[T, E]) Ok(value T) Result[T, E] {
	return NewOk[T, E](value)
}
func (r *RH[T, E]) Err(err E) Result[T, E] {
	return NewErr[T, E](err)
}
