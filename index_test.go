package goaction

import "testing"

func TestReturnStr(t *testing.T) {
type args struct {
	str string
}
tests := []struct {
	name string
	args args
	want string
}{
	{
		name: "Should Return same",
		args: args{"Hello"},
		want: "Hello",
	},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReturnStr(tt.args.str); got != tt.want {
				t.Errorf("ReturnStr() = %v, want %v", got, tt.want)
			}
		})
	}
}