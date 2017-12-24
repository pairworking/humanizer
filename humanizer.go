package humanize

type Humanize interface {
	ToWords(interface{}) (string, error)
}

type Validator interface {
	IsValid(interface{}) bool
}
