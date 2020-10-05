package arrays

import "testing"

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
		{"with valid test case", args{[]int{2,1,5,3,4}}, 3, false},
		{"with unsolvable", args{[]int{2,5,1,3,4}}, 0, true},
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
