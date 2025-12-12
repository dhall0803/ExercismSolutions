package protein

import (
	"errors"
	"slices"
)

var ErrStop = errors.New("stop error")
var ErrInvalidBase = errors.New("invalid base error")

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

var stopCodes = []string{
	"UAA",
	"UAG",
	"UGA",
}

const codonLength = 3

func rnaToCodons(rna string) ([]string, error) {
	if len(rna)%codonLength != 0 {
		return []string{}, ErrInvalidBase
	}
	result := []string{}
	for i := 0; i < len(rna); i += 3 {
		result = append(result, rna[i:i+3])
	}
	return result, nil
}

func FromRNA(rna string) ([]string, error) {
	codons, err := rnaToCodons(rna)
	if err != nil {
		return []string{}, err
	}
	result := []string{}
	for _, codon := range codons {
		aminoAcid, err := FromCodon(codon)
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return []string{}, err
		}
		result = append(result, aminoAcid)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	if slices.Contains(stopCodes, codon) {
		return "", ErrStop
	}

	if val, ok := codonToProtein[codon]; ok {
		return val, nil
	}

	return "", ErrInvalidBase
}
