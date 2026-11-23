package ascii

import (
	"strings"
	"testing"
)

func TestTransformText(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		banner  string
		wantErr bool
	}{
		{
			name:    "Valid text and banner",
			text:    "Hello",
			banner:  "../banner/standard.txt",
			wantErr: false,
		},
		{
			name:    "Invalid banner",
			text:    "Hello",
			banner:  "../banner/invalid.txt",
			wantErr: true,
		},
		{
			name:    "Text with newline literal",
			text:    "Hello\\nWorld",
			banner:  "../banner/standard.txt",
			wantErr: false,
		},
		{
			name:    "Empty string",
			text:    "",
			banner:  "../banner/standard.txt",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformText(tt.text, tt.banner)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.text == "" {
				if got != "\n" {
					t.Errorf("TransformText() for empty string should return a newline, got %q", got)
				}
			}
			if !tt.wantErr && tt.text == "Hello" {
				if !strings.Contains(got, "H") && len(got) < 10 { // simplistic check
					t.Errorf("TransformText() result seems incorrect, got length %d", len(got))
				}
			}
		})
	}
}

func TestReadBanner(t *testing.T) {
	tests := []struct {
		name    string
		txtfile string
		wantErr bool
	}{
		{
			name:    "Valid banner file",
			txtfile: "../banner/standard.txt",
			wantErr: false,
		},
		{
			name:    "Invalid banner file",
			txtfile: "../banner/nonexistent.txt",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := ReadBanner(tt.txtfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBanner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(lines) == 0 {
				t.Errorf("ReadBanner() returned empty slice for valid file")
			}
		})
	}
}
