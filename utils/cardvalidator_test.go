package utils

import (
	"testing"
)

func TestCardValidator_FirstNumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		cv   *CardValidator
		args args
		want string
	}{
		{
			name: "Positive: first char number 4",
			cv:   &CardValidator{},
			args: args{
				value: "41234567890123456789",
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &CardValidator{}
			if got := cv.FirstNumber(tt.args.value); got != tt.want {
				t.Errorf("CardValidator.FirstNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardValidator_InRange(t *testing.T) {
	v := "5233111222211112222"
	vv := v[0:4]
	type args struct {
		value string
		lower string
		upper string
	}
	tests := []struct {
		name string
		cv   *CardValidator
		args args
		want bool
	}{
		{
			name: "Positive: 5233111222211112222",
			cv:   &CardValidator{},
			args: args{
				value: vv,
				lower: "2221",
				upper: "2720",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &CardValidator{}
			if got := cv.InRange(tt.args.value, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("CardValidator.InRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
