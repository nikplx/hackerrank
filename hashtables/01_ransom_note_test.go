package hashtables

import "testing"

func Test_checkMagazine(t *testing.T) {
	type args struct {
		magazine []string
		note     []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"with one word missing", args{
			magazine: []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
			note: []string{"ive", "got", "some", "coconuts"}},
			false,
		},
		{"with all words in", args{
			magazine: []string{"two", "times", "two", "is", "not", "four"},
			note: []string{"two", "times", "two", "is",
		"four"}},
			true,
		},
		{"with all words in but one too little", args{
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note: []string{"two", "times", "two", "is",
		"four"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMagazine(tt.args.magazine, tt.args.note); got != tt.want {
				t.Errorf("checkMagazine() = %v, want %v", got, tt.want)
			}
		})
	}
}
