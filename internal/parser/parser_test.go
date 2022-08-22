package parser

import "testing"

func TestParser_Preprocess(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should successfully return preprocessed word",
			args: args{word: "~~Hello~~"},
			want: "hello",
		},
		{
			name: "should return empty string when no regex match",
			args: args{word: "~~!!09!~~"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{}
			if got := p.Preprocess(tt.args.word); got != tt.want {
				t.Errorf("Preprocess() = %v, want %v", got, tt.want)
			}
		})
	}
}
