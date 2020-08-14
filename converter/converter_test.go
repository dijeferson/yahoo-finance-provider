package converter

import "testing"

func TestConvertToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{
			name:  "Case Integer",
			input: "10",
			want:  10,
		},
		{
			name:  "Case Float",
			input: "3.1415",
			want:  3.1415,
		},
		{
			name:  "Case not parseable",
			input: "ABC",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToFloat64(tt.input); got != tt.want {
				t.Errorf("ConvertToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
