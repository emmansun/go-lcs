package main

import (
	"log"

	"github.com/emmansun/go-lcs"
)

func convert(s string) []interface{} {
	r := []rune(s)
	result := make([]interface{}, len(r))
	for i, v := range r {
		result[i] = v
	}
	return result
}

func main() {
	fastLCS := lcs.NewFastLCS(convert("GAC"), convert("AGCAT"))
	pairs := fastLCS.FindAllLcsPairs()
	log.Printf("Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCS(convert("BADCDCBA"), convert("ABCDCDAB"))
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCS(convert("ABAB"), convert("ABABAB"))
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCS(convert("ABCAAABBABBCCABCBACABABABCCBC"), convert("ABCABABABCBACABCBACABABACBCB"))
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)
}
