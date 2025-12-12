package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var result []string
	for i := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			break
		}

		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	return result
}
