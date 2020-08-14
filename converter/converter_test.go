package converter

import "testing"

func TestConvertToFloat64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToFloat64(tt.args.value); got != tt.want {
				t.Errorf("ConvertToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
