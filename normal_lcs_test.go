package lcs

import (
	"reflect"
	"testing"
)

func TestNormalLCS_FindAllLcsPairs(t *testing.T) {
	tests := []struct {
		name string
		this *NormalLCS
		want int
	}{
		// TODO: Add test cases.
		{"wiki_sample", NewNormalLCSString("GAC", "AGCAT"), 3},
		{"sample", NewNormalLCSString("BADCDCBA", "ABCDCDAB"), 8},
		{"loop1", NewNormalLCSString("ABAB", "ABABAB"), 5},
		{"complexCase", NewNormalLCSString("ABCAAABBABBCCABCBACABABABCCBC", "ABCABABABCBACABCBACABABACBCB"), 220},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(tt.this.FindAllLcsPairs()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalLCS.FindAllLcsPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
