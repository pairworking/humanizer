package humanize

import (
	"bytes"
	"errors"
)

type NumberValidator struct{}

func (NumberValidator) isValid(n interface{}) bool {
	switch t := n.(type) {
	case int, int32, int64:
		_ = t
		return true
	default:
		return false
	}
}

type Humanizer struct {
}

func (n *Humanizer) ToWords(item interface{}) (string, error) {

	numberValidator := &NumberValidator{}
	if !numberValidator.isValid(item) {
		return "", ErrUnsuportedDataType
	}

	return "", nil

}

type NumberHumanizer struct {
}

const (
	Ten         = 10
	Hundred     = 100
	Thousand    = 1000
	Million     = 1000000
	Billion     = 1000000000
	Trillion    = 1000000000000
	Quadrillion = 1000000000000000
	Quintillion = 1000000000000000000
	Sextillion  = 1000000000000000000000
	Septillion  = 1000000000000000000000000
	Octillion   = 1000000000000000000000000000
	Nonillion   = 1000000000000000000000000000000
)

var (
	ErrUnsuportedDataType = errors.New("Data type is not supported yet")
)

func (n *NumberHumanizer) ToWords(number int64) (string, error) {

	res := n.divide(number)
	return res, nil
}

func (n *NumberHumanizer) divide(number int64) string {

	value := number
	bufferRes := bytes.NewBufferString("")
	if value == 0 {
		return "zero"
	}

	if value < 0 {
		bufferRes.WriteString("minus")
		value *= -1
		bufferRes.WriteString(n.divide(value))
		return bufferRes.String()
	}

	if temp := value / Quintillion; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" quintillion")
		value = value % Quintillion
	}

	if temp := value / Quadrillion; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" quadrillion")
		value = value % Quadrillion
	}

	if temp := value / Trillion; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" trillion")
		value = value % Trillion
	}

	if temp := value / Billion; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" billion")
		value = value % Billion
	}

	if temp := value / Million; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" million")
		value = value % Million
	}

	if temp := value / Thousand; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" thousand")
		value = value % Thousand
	}

	if temp := value / Hundred; temp > 0 {
		bufferRes.WriteString(n.divide(temp))
		bufferRes.WriteString(" hundred")
		value = value % Hundred
	}
	if value > 0 {
		if bufferRes.Len() > 0 {
			bufferRes.WriteString(" and")
		}
		bufferRes.WriteString(" ")
		if value < 20 {
			bufferRes.WriteString(unitArr[value])
		} else {
			tenVal := tensArr[value/10]
			bufferRes.WriteString(tenVal)
			if mod := value % 10; mod > 0 {
				bufferRes.WriteString("-")
				bufferRes.WriteString(unitArr[mod])
			}
		}
	}

	return bufferRes.String()
}

var (
	unitArr = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}
	tensArr = []string{
		"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}
)

func NumberToWords(number int64) string {
	numHumanizer := NumberHumanizer{}
	res, _ := numHumanizer.ToWords(number)
	return res
}
