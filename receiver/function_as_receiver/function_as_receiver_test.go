package receiver

import "testing"

func TestF_upperCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		f    F
		args args
		want string
	}{
		{
			name: "upperDoubleStringcase",
			f:    F(doubleString),
			args: args{
				s: "hello",
			},
			want: "HELLOHELLO",
		},
		{
			name: "upperinitString",
			f:    F(initial),
			args: args{
				s: "hello",
			},
			want: "H",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.upperCase(tt.args.s); got != tt.want {
				t.Errorf("F.upperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
