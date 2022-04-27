package util

import (
	"github.com/nahida05/query-monitoring/internal/storage/model"
	"testing"
)

func TestExists(t *testing.T) {
	testCases := []struct {
		value    string
		options  []string
		expected bool
	}{{
		value:    "select",
		options:  []string{model.SELECT, model.INSERT},
		expected: true,
	},
		{
			value:    "create",
			options:  []string{model.SELECT, model.INSERT, model.UPDATE},
			expected: false,
		},
		{
			value:    "SELECT",
			options:  []string{model.SELECT, model.INSERT},
			expected: true,
		},
		{
			value:    "select",
			options:  nil,
			expected: false,
		},
	}

	for _, test := range testCases {
		if got := Exists(test.value, test.options...); test.expected != got {
			t.Errorf("expected %t got %t", test.expected, got)

		}
	}
}

func TestPageCount(t *testing.T) {
	type args struct {
		total        int
		limit        int
		defaultLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with remain number",
			args: args{
				total:        23,
				limit:        10,
				defaultLimit: 10,
			},
			want: 3,
		},
		{
			name: "with valid number",
			args: args{
				total:        45,
				limit:        5,
				defaultLimit: 10,
			},
			want: 9,
		},
		{
			name: "with zero limit",
			args: args{
				total:        41,
				limit:        0,
				defaultLimit: 10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageCount(tt.args.total, tt.args.limit, tt.args.defaultLimit); got != tt.want {
				t.Errorf("PageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
