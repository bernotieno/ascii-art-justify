package Ascii

import (
	"testing"
)

func TestHandleSpecialCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "No special cases",
			args:    args{s: "Normal string"},
			wantErr: false,
		},
		{
			name:    "Contains \\a",
			args:    args{s: "String with \\a"},
			wantErr: true,
		},
		{
			name:    "Contains \\b",
			args:    args{s: "String with \\b"},
			wantErr: true,
		},
		{
			name:    "Contains \\t",
			args:    args{s: "String with \\t"},
			wantErr: true,
		},
		{
			name:    "Contains \\f",
			args:    args{s: "String with \\f"},
			wantErr: true,
		},
		{
			name:    "Contains \\r",
			args:    args{s: "String with \\r"},
			wantErr: true,
		},
		{
			name:    "Contains \\v",
			args:    args{s: "String with \\v"},
			wantErr: true,
		},
		{
			name:    "Contains multiple special cases",
			args:    args{s: "String with \\a and \\t"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := HandleSpecialCase(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleSpecialCase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
