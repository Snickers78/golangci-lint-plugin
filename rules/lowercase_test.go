package rules

import "testing"

func TestIsLowercase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "starts with uppercase",
			s:    "Starting server on port 8080",
			want: true,
		},
		{
			name: "starts with lowercase",
			s:    "starting server on port 8080",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isLowercase(tt.s)
			if got != tt.want {
				t.Fatalf("containsNonEnglishLetters(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
