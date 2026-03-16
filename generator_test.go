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
			name:    "configuración válida por defecto",
			opts:    Options{Length: 12, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: false,
		},
		{
			name:    "solo minúsculas",
			opts:    Options{Length: 8, UseLower: true},
			wantErr: false,
		},
		{
			name:    "longitud máxima permitida",
			opts:    Options{Length: 128, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: false,
		},
		{
			name:    "longitud cero",
			opts:    Options{Length: 0, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "longitud negativa",
			opts:    Options{Length: -5, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "excede longitud máxima",
			opts:    Options{Length: 129, UseUpper: true, UseLower: true, UseNumbers: true, UseSymbols: true},
			wantErr: true,
		},
		{
			name:    "sin tipos de caracteres seleccionados",
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
				t.Errorf("Generate() longitud = %d, quería %d", len(got), tt.opts.Length)
			}
		})
	}
}

func TestGenerateCharset(t *testing.T) {
	t.Run("solo mayúsculas", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseUpper: true})
		if err != nil {
			t.Fatal(err)
		}
		if strings.ToUpper(pass) != pass {
			t.Errorf("se esperaban solo mayúsculas, got: %q", pass)
		}
	})

	t.Run("solo dígitos", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseNumbers: true})
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range pass {
			if c < '0' || c > '9' {
				t.Errorf("carácter inesperado fuera de dígitos: %c", c)
			}
		}
	})

	t.Run("solo minúsculas", func(t *testing.T) {
		pass, err := Generate(Options{Length: 30, UseLower: true})
		if err != nil {
			t.Fatal(err)
		}
		if strings.ToLower(pass) != pass {
			t.Errorf("se esperaban solo minúsculas, got: %q", pass)
		}
	})
}

func TestStrengthScore(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{"contraseña muy corta", "abc", "Débil"},
		{"solo minúsculas sin longitud suficiente", "abcdefg", "Débil"},
		{"longitud media con variedad", "Abcdef12", "Media"},
		{"larga pero sin variedad", "aaaaaaaaaaaaaaaaaa", "Media"},
		{"buena variedad y longitud", "Abcdef12!", "Fuerte"},
		{"muy larga con todo", "Abcdef12!@#XYZabc1", "Muy fuerte"},
		{"muy fuerte con longitud y variedad", "Abcdef12!@#XYZ", "Muy fuerte"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrengthScore(tt.password)
			if got != tt.want {
				t.Errorf("StrengthScore(%q) = %q, quería %q", tt.password, got, tt.want)
			}
		})
	}
}
