package rules

import "testing"

func TestContainsSpecialSymbolsOrEmoji(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "letters and spaces",
			s:    "server started successfully",
			want: false,
		},
		{
			name: "letters digits and spaces",
			s:    "server started on port 8080",
			want: false,
		},
		{
			name: "comma is special symbol",
			s:    "server started, ok",
			want: true,
		},
		{
			name: "emoji is special symbol",
			s:    "server started 🚀",
			want: true,
		},
		{
			name: "hash sign is special symbol",
			s:    "error #123",
			want: true,
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

			got := containsSpecialSymbolsOrEmoji(tt.s)
			if got != tt.want {
				t.Fatalf("containsSpecialSymbolsOrEmoji(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
