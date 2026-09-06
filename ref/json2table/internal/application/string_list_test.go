package application

import (
	"reflect"
	"testing"
)

func TestStringListUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    StringList
		wantErr bool
	}{
		{
			name:    "single string",
			input:   `"alpha"`,
			want:    StringList{"alpha"},
			wantErr: false,
		},
		{
			name:    "string array",
			input:   `["alpha","beta"]`,
			want:    StringList{"alpha", "beta"},
			wantErr: false,
		},
		{
			name:    "empty string array",
			input:   `[]`,
			want:    StringList{},
			wantErr: false,
		},
		{
			name:    "number is invalid",
			input:   `123`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "malformed json is invalid",
			input:   `["alpha"`,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got StringList
			err := got.UnmarshalJSON([]byte(tc.input))

			if tc.wantErr {
				if err == nil {
					t.Fatalf("UnmarshalJSON(%s) expected error, got nil", tc.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("UnmarshalJSON(%s) returned error: %v", tc.input, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("UnmarshalJSON(%s) = %#v, want %#v", tc.input, got, tc.want)
			}
		})
	}
}
