package humanize

import "fmt"

func ExampleNumber() {

	res := NumberToWords(1335168736797130752)
	fmt.Println(res)
	// Output: one quintillion three hundred and thirty-five quadrillion one hundred and sixty-eight trillion seven hundred and thirty-six billion seven hundred and ninety-seven million one hundred and thirty thousand seven hundred and fifty-two
}
