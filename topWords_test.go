package hw4

import (
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T) {
	testTable := []struct{
		arg1 string
		arg2 int
		want []string

	}{
		{
			arg1: "amogus is a bad game, amogus, truly a bad game, amogus, bad, amogus.",
			arg2: 1,
			want: []string{"amogus"},
		},
		{
			arg1: "amogus is a bad game, amogus, truly a bad game, amogus, bad.",
			arg2: 2,
			want: []string{"amogus", "bad", "a", "game"},
		},
		{
			arg1: "one two three two three four three four four four",
			arg2: 3,
			want: []string{"four", "three", "two"},
		},
		{
			arg1: "one two three two three four three four four four",
			arg2: 2,
			want: []string{"four", "three"},
		},
	}
	for _, v := range testTable {
		got := topWords(v.arg1, v.arg2)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("Result of topWords was incorrect, got: %v, want: %v\n", got, v.want)
		}
	}
}
