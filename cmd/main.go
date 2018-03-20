package main

import (
	"log"

	"github.com/emmansun/go-lcs"
)

func main() {

	fastLCS := lcs.NewFastLCSString("GAC", "AGCAT")
	pairs := fastLCS.FindAllLcsPairs()
	log.Printf("[FAST] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	normalLCS := lcs.NewNormalLCSString("GAC", "AGCAT")
	pairs = normalLCS.FindAllLcsPairs()
	log.Printf("[NORMAL] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCSString("BADCDCBA", "ABCDCDAB")
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("[FAST] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	normalLCS = lcs.NewNormalLCSString("BADCDCBA", "ABCDCDAB")
	pairs = normalLCS.FindAllLcsPairs()
	log.Printf("[NORMAL] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCSString("ABAB", "ABABAB")
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("[FAST] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	normalLCS = lcs.NewNormalLCSString("ABAB", "ABABAB")
	pairs = normalLCS.FindAllLcsPairs()
	log.Printf("[NORMAL] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	fastLCS = lcs.NewFastLCSString("ABCAAABBABBCCABCBACABABABCCBC", "ABCABABABCBACABCBACABABACBCB")
	pairs = fastLCS.FindAllLcsPairs()
	log.Printf("[FAST] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)

	normalLCS = lcs.NewNormalLCSString("ABCAAABBABBCCABCBACABABABCCBC", "ABCABABABCBACABCBACABABACBCB")
	pairs = normalLCS.FindAllLcsPairs()
	log.Printf("[NORMAL] Candidate size=%d, Candicates=%v\n\n", len(pairs), pairs)
}
