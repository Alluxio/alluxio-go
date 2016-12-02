package option

type Exists struct{}

func RandomExists() Exists {
	var option Exists
	return option
}
