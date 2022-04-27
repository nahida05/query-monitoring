package service

import (
	"reflect"
	"testing"
)

func Test_newCustomError(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want customError
	}{
		{
			name: "with empty value",
			args: args{content: ""},
			want: customError{},
		},
		{
			name: "with non empty value",
			args: args{content: "bad request"},
			want: customError{Message: "bad request"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCustomError(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCustomError() = %v, want %v", got, tt.want)
			}
		})
	}
}
