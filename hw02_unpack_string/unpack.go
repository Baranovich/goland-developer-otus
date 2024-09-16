package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sequence string) (string, error) {
	sequenceRunes := []rune(sequence)
	var transformedSequence strings.Builder
	basket := ""
	for i := 0; i < len(sequenceRunes); i++ {
		symbol, err := getSymbolFromSequence(sequenceRunes, &i)
		if err != nil {
			return "", ErrInvalidString
		}
		if number, err := strconv.Atoi(symbol); err == nil {
			if basket == "" {
				return "", ErrInvalidString
			}
			symbolToWrite, err := getSymbolToWriteTransformedSequence(basket)
			if err != nil {
				return "", ErrInvalidString
			}
			transformedSequence.WriteString(strings.Repeat(symbolToWrite, number))
			basket = ""
		} else {
			err := emptyTheBasketIfContainsSmth(&basket, &transformedSequence)
			if err != nil {
				return "", ErrInvalidString
			}
			basket = symbol
		}
	}
	err := emptyTheBasketIfContainsSmth(&basket, &transformedSequence)
	if err != nil {
		return "", ErrInvalidString
	}
	return transformedSequence.String(), nil
}

func emptyTheBasketIfContainsSmth(basketPointer *string, transformedSequencePointer *strings.Builder) error {
	if *basketPointer != "" {
		symbolToWrite, err := getSymbolToWriteTransformedSequence(*basketPointer)
		if err != nil {
			return ErrInvalidString
		}
		(*transformedSequencePointer).WriteString(symbolToWrite)
	}
	return nil
}

func getSymbolFromSequence(sequence []rune, i *int) (string, error) {
	var symbol string
	if string(sequence[*i]) == `\` {
		if *i+1 >= len(sequence) {
			return "", ErrInvalidString
		}
		symbol = string(sequence[*i]) + string(sequence[*i+1])
		*i++
	} else {
		symbol = string(sequence[*i])
	}
	return symbol, nil
}

func getSymbolToWriteTransformedSequence(symbol string) (string, error) {
	if len(symbol) == 2 {
		switch symbol {
		case `\\`, `\0`, `\1`, `\2`, `\3`, `\4`, `\5`, `\6`, `\7`, `\8`, `\9`:
			return string([]rune(symbol)[1]), nil
		}
		return "", ErrInvalidString
	}
	return symbol, nil
}
