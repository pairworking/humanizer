package humanizer

type Humanize interface {
	ToWords(interface{}) (string, error)
}

type Validator interface {
	IsValid(interface{}) bool
}
