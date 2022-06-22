package gtintools

import (
	"testing"
)

func TestValidateGtinInvalidCharacters(t *testing.T) {
	gtins := [2]string{"a2345670", "1234567a"}

	t.Log("Test gtin contains only digits validation function")
	{
		for _, gtin := range gtins {
			isValid := gtinContainsOnlyDigits(gtin)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
		}
	}

	t.Log("Test global validation function")
	{
		for _, gtin := range gtins {
			isValid, err := ValidateGtin(gtin)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
			if err != ErrInvalidCharacters {
				t.Errorf("ErrInvalidCharacters expected for GTIN '%v'", gtin)
			}
		}
	}
}

func TestValidateGtinInvalidLength(t *testing.T) {
	gtins := [4]string{"1234567", "123456789", "12345678901", "123456789012345"}

	t.Log("Test length validation function")
	{
		for _, gtin := range gtins {
			isValid := gtinLengthIsValid(gtin)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
		}
	}

	t.Log("Test global Validation function")
	{
		for _, gtin := range gtins {
			isValid, err := ValidateGtin(gtin)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
			if err != ErrInvalidLength {
				t.Errorf("ErrInvalidLength expected for GTIN '%v'", gtin)
			}
		}
	}
}

func TestValidateGtinInvalidCheckDigit(t *testing.T) {
	gtins := [7]string{"12345678", "00000012345678", "123456789019", "00123456789019", "1234567890123", "01234567890123", "12345678901234"}

	t.Log("Test check digit validation function")
	{
		for _, gtin := range gtins {
			isValid := checkDigitIsCorrect(gtin, 0)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
		}
	}

	t.Log("Test global Validation function")
	{
		for _, gtin := range gtins {
			isValid, err := ValidateGtin(gtin)
			if isValid == true {
				t.Errorf("Format of GTIN '%v' should be invalid", gtin)
			}
			if err != ErrInvalidCheckDigit {
				t.Errorf("ErrInvalidCheckDigit expected for GTIN '%v'", gtin)
			}
		}
	}
}

func TestValidateGtinValidGtin(t *testing.T) {
	gtins := [7]string{"12345670", "00000012345670", "123456789012", "00123456789012", "1234567890128", "01234567890128", "12345678901231"}

	for _, gtin := range gtins {
		isValid, err := ValidateGtin(gtin)
		if isValid == false {
			t.Errorf("Format of GTIN '%v' should be valid", gtin)
		}
		if err != nil {
			t.Errorf("No error for GTIN '%v' should be returned", gtin)
		}
	}
}
