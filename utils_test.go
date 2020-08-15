package yahoofinance

import (
	"reflect"
	"testing"
)

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
			if got := convertToFloat64(tt.input); got != tt.want {
				t.Errorf("ConvertToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCSVFromString(t *testing.T) {
	type args struct {
		contents     string
		removeHeader bool
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Sucess no header",
			args: args{contents: "A,B", removeHeader: false},
			want: [][]string{
				{"A", "B"},
			},
			wantErr: false,
		},
		{
			name: "Sucess with header",
			args: args{contents: "field1,field2\nA,B", removeHeader: true},
			want: [][]string{
				{"A", "B"},
			},
			wantErr: false,
		},
		{
			name:    "Sucess with empty input and removeHeader",
			args:    args{contents: "", removeHeader: true},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Sucess with empty input and dont removeHeader",
			args:    args{contents: "", removeHeader: false},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Sucess with removeHeader when no header",
			args:    args{contents: "A,B", removeHeader: true},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Error on bad number of columns",
			args:    args{contents: "A,B\nC,D,E", removeHeader: true},
			want:    [][]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCSVFromString(tt.args.contents, tt.args.removeHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSVFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
