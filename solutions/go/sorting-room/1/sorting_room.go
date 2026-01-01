package sorting


import (
	"fmt"
	"strconv"
)

type NumberBox interface {
	Number() int
}

type FancyNumberBox interface {
	Value() string
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	floatNumber := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", floatNumber)
}

// ExtractFancyNumber should return the integer value for a FancyNumber
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if nb, ok := fnb.(FancyNumber); ok {
		if n, err := strconv.Atoi(nb.Value()); err == nil {
			return n
		}
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if nb, ok := fnb.(FancyNumber); ok {
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(nb)))
	}

	return "This is a fancy box containing the number 0.0"
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case float64:
		return DescribeNumber(v)
	case int:
		return DescribeNumber(float64(v))
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}

