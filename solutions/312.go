package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findProvider(card string) string{
	prefixes := map[string]string{
		"4":  "Visa",
		"51": "MasterCard",
		"52": "MasterCard",
		"53": "MasterCard",
		"54": "MasterCard",
		"55": "MasterCard",
		"34": "Amex",
		"37": "Amex",
		"50": "Maestro",
	}

	for prefix, provider := range prefixes{
		if strings.HasPrefix(card, prefix) {
			return provider
		}
	}
	return ""
}

func luhnCheck(card string) (bool, string) {
	sum := 0
	chars := strings.Split(card, "")
	count := len(chars)

	if count < 2 {
		return false, ""
	}

	for i := count - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(chars[i])
		doubled := 0
		if (count - i - 1) % 2 == 0 {
			doubled := digit * 2
			if doubled > 9 {
				doubled -= 9
			}
		}
		sum += doubled
	}

	isValid := sum % 10 == 0
	provider := findProvider(card)
	return isValid, provider
}

func main() {
	isValid, provider := luhnCheck("5555555555554444")
	fmt.Println(isValid)
	fmt.Println(provider)
}
