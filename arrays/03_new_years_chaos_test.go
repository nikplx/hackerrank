package arrays

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"testing"
)

func Test_minimumBribes(t *testing.T) {
	type args struct {
		queue []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"from hackkerrank samples", args{[]int{2,1,5,3,4}}, 3, false},
		{"with 1 swap", args{[]int{2,1,3,4}}, 1, false},
		{"with 0 swaps", args{[]int{1,2,3}}, 0, false},
		{"with too chaotic", args{[]int{2,5,1,3,4}}, 0, true},
		{"with sorted duplicates", args{[]int{2,5,5,6}}, 0, false},
		{"with unsorted duplicates", args{[]int{1,5,6,5}}, 1, false},
		{"with very big shuffeled list", args{getVeryBigIntList(t, "very_big_list_unsorted.json")}, 0, true},
		{"with very big valid list", args{getVeryBigIntList(t, "very_big_list_sorted.json")}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minimumBribes(tt.args.queue)
			if (err != nil) != tt.wantErr {
				t.Errorf("minimumBribes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("minimumBribes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getVeryBigIntList(t *testing.T, name string) []int {
	body, err := ioutil.ReadFile(path.Join("test_data", name))
	if err != nil {
		t.Fatal(err)
	}

	var res []int
	if err := json.Unmarshal(body, &res); err != nil {
		t.Fatal(err)
	}

	return res
}
