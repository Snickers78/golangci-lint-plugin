package rules

import "testing"

func TestContainsSensitiveData_Keywords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "password keyword",
			s:    "user password is here",
			want: false,
		},
		{
			name: "token keyword upper case",
			s:    "USER TOKEN",
			want: false,
		},
		{
			name: "no sensitive keywords",
			s:    "user logged in successfully",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _ := containsSensitiveData(tt.s)
			if got != tt.want {
				t.Fatalf("containsSensitiveData(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestContainsSensitiveData_Patterns(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "email address",
			s:    "contact me at test@example.com",
			want: true,
		},
		{
			name: "credit card like number",
			s:    "card number 4111 1111 1111 1111",
			want: true,
		},
		{
			name: "ip address",
			s:    "request from 10.0.0.1 failed",
			want: true,
		},
		{
			name: "safe message",
			s:    "server started on port 8080",
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _ := containsSensitiveData(tt.s)
			if got != tt.want {
				t.Fatalf("containsSensitiveData(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
