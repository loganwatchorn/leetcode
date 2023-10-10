package main

func reverseWords(s string) string {
	res := make([]byte, 0, len(s))
	word := []byte{}

	for i := range s {
		if s[i] == ' ' {
			for i := len(word) - 1; i >= 0; i-- {
				res = append(res, word[i])
			}
			word = []byte{}
			res = append(res, ' ')
			continue
		}

		word = append(word, s[i])
	}

	for i := len(word) - 1; i >= 0; i-- {
		res = append(res, word[i])
	}

	return string(res)
}
