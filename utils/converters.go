package utils

import "strconv"

/**
*THIS PACKAGE DEALS with type conversion from num and float to string
 */

func ChangeFloattoString(input float64) string {
	return strconv.FormatFloat(input, 'f', -1, 32)
}
