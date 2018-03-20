package lcs

import (
	"reflect"
	"testing"
)

func TestFastLCS_FindAllLcsPairs(t *testing.T) {
	tests := []struct {
		name string
		this *FastLCS
		want int
	}{
		// TODO: Add test cases.
		{"wiki_sample", NewFastLCSString("GAC", "AGCAT"), 3},
		{"sample", NewFastLCSString("BADCDCBA", "ABCDCDAB"), 8},
		{"loop1", NewFastLCSString("ABAB", "ABABAB"), 5},
		{"complexCase", NewFastLCSString("ABCAAABBABBCCABCBACABABABCCBC", "ABCABABABCBACABCBACABABACBCB"), 220},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(tt.this.FindAllLcsPairs()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FastLCS.FindAllLcsPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
