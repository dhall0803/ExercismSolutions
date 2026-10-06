package acronym

import(
    "strings"
    "unicode"
)

func hypenToSpace(s string) string {
    result := []rune{}
    
    for _, r := range s {
        if r == '-' {
            result = append(result, ' ')
            continue
        }
        result = append(result, r)
    }
    
    return string(result)
}

func sanitise(s string) string {
    result := []rune{}

    for _, r := range s {
        if !unicode.IsLetter(r) {
            continue
        }

        result = append(result, r)
    }

    return string(result)
}

func sanitiseWords(words []string) []string {
    sanitisedWords := []string{}

    for _, word := range words {
        sanitisedWords = append(sanitisedWords, sanitise(word))
    }

    return sanitisedWords
}

func Abbreviate(s string) string {
	normalisedString := hypenToSpace(s)
	words := strings.Fields(normalisedString)

	sanitisedWords := sanitiseWords(words)
    
    letters := []string{}
    for _, word := range sanitisedWords{
        letters = append(letters, string(strings.ToUpper(word)[0]))        
    }
	return strings.Join(letters, "")
}
