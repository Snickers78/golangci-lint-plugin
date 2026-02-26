package rules

import "testing"

func TestContainsNonEnglishLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "english only",
			s:    "hello world",
			want: false,
		},
		{
			name: "russian only",
			s:    "привет мир",
			want: true,
		},
		{
			name: "mixed english and russian",
			s:    "hello мир",
			want: true,
		},
		{
			name: "digits only",
			s:    "123456",
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

			got := containsNonEnglishLetters(tt.s)
			if got != tt.want {
				t.Fatalf("containsNonEnglishLetters(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
