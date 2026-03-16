package generator

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		opts    Options
		wantErr bool
	}{
		{
			name:    "valid default configuration",
			opts:    Options{Length: 12, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: false,
		},
		{
			name:    "lowercase only",
			opts:    Options{Length: 8, UseLower: true},
			wantErr: false,
		},
		{
			name:    "maximum allowed length",
			opts:    Options{Length: 128, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: false,
		},
		{
			name:    "zero length",
			opts:    Options{Length: 0, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "negative length",
			opts:    Options{Length: -5, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "exceeds maximum length",
			opts:    Options{Length: 129, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "no character types selected",
			opts:    Options{Length: 12},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.opts.Length {
				t.Errorf("Generate() length = %d, want %d", len(got), tt.opts.Length)
			}
		})
	}
}

func TestGenerateCharset(t *testing.T) {
	t.Run("uppercase only", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseUpper: true})
		if err != nil {
			t.Fatal(err)
		}
		if strings.ToUpper(pass) != pass {
			t.Errorf("expected uppercase only, got: %q", pass)
		}
	})

	t.Run("digits only", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseNumbers: true})
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range pass {
			if c < '0' || c > '9' {
				t.Errorf("unexpected character outside digits: %c", c)
			}
		}
	})

	t.Run("lowercase only", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseLower: true})
		if err != nil {
			t.Fatal(err)
		}
		if strings.ToLower(pass) != pass {
			t.Errorf("expected lowercase only, got: %q", pass)
		}
	})
}

func TestStrengthScore(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{"very short password", "abc", "Weak"},
		{"lowercase only without sufficient length", "abcdefg", "Weak"},
		{"medium length with variety", "Abcdef12", "Medium"},
		{"long but no variety", "aaaaaaaaaaaaaaaaaa", "Medium"},
		{"good variety and length", "Abcdef12!", "Strong"},
		{"very long with everything", "Abcdef12!@#XYZabc1", "Very strong"},
		{"very strong with length and variety", "Abcdef12!@#XYZ", "Very strong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrengthScore(tt.password)
			if got != tt.want {
				t.Errorf("StrengthScore(%q) = %q, want %q", tt.password, got, tt.want)
			}
		})
	}
}
