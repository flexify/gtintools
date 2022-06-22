package gtintools

import (
	"errors"
	"math"
	"regexp"
)

var (
	ErrInvalidCharacters = errors.New("only digits are allowed")
	ErrInvalidLength     = errors.New("GTINs must be 8, 12, 13, or 14 digits long")
	ErrInvalidCheckDigit = errors.New("check digit is not correct")
)

// ValidateGtin
// the GTIN is walid when
// - only digits are given
// - the length is 8, 12, 13, or 14
// - the check digit is correct
func ValidateGtin(gtin string) (isValid bool, err error) {

	// Validate that only digits are given
	if !gtinContainsOnlyDigits(gtin) {
		return false, ErrInvalidCharacters
	}

	// Validate the length
	if !gtinLengthIsValid(gtin) {
		return false, ErrInvalidLength
	}

	// Validate the check digit
	checkDigit := calculateCheckDigit(gtin)
	if !checkDigitIsCorrect(gtin, checkDigit) {
		return false, ErrInvalidCheckDigit
	}

	return true, nil
}

func gtinContainsOnlyDigits(gtin string) bool {
	onlyDigitsRegExp, _ := regexp.Compile(`^\d+$`)
	return onlyDigitsRegExp.MatchString(gtin)
}

func gtinLengthIsValid(gtin string) bool {
	lengthRegExp, _ := regexp.Compile("^(?:.{8}|.{12}|.{13}|.{14})$")
	return lengthRegExp.MatchString(gtin)
}

func calculateCheckDigit(gtin string) byte {
	checksum := float64(calculateChecksum(gtin))

	// Calculate the check digit
	return byte((math.Ceil(checksum/10) * 10) - checksum)
}

func calculateChecksum(gtin string) uint8 {
	// Build slice with all numbers except check digit with leading zeros
	gtinLength := len(gtin)
	numbersWithoutCheckDigit := make([]uint8, 14-gtinLength)
	for i := 0; i < gtinLength-1; i++ {
		number := uint8(gtin[i] - '0')
		numbersWithoutCheckDigit = append(numbersWithoutCheckDigit, number)
	}

	// Calculate the checksum
	var checksum uint8
	for i, v := range numbersWithoutCheckDigit {
		if i%2 == 0 {
			checksum += v * 3
		} else {
			checksum += v
		}
	}

	return checksum
}

func checkDigitIsCorrect(gtin string, checkDigit byte) bool {
	var givenCheckDigit = gtin[len(gtin)-1] - '0'
	return checkDigit == givenCheckDigit
}
