package palindromic_substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSubstrings(tt.args.s); got != tt.want {
				t.Errorf("CountSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
