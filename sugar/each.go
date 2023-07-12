package sugar

// MapEach applies a function to each item in a list and returns a new list
func MapEach[T any](list []T, f func(t T) T) []T {
	result := make([]T, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	return result
}

func ForEeach[T any](list []T, f func(t T) error) error {
	for _, item := range list {
		err := f(item)
		if err != nil {
			return err
		}
	}
	return nil
}
