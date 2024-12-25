package cli

import "testing"

// Тест для функции checkbox
func TestCheckbox(t *testing.T) {
	tests := []struct {
		label   string
		checked bool
		want    string
	}{
		{"Option 1", true, "[x] Option 1"},
		{"Option 2", false, "[ ] Option 2"},
		{"Option 3", true, "[x] Option 3"},
		{"Option 4", false, "[ ] Option 4"},
	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			got := checkbox(tt.label, tt.checked)
			if got != tt.want {
				t.Errorf("checkbox(%q, %v) = %q; want %q", tt.label, tt.checked, got, tt.want)
			}
		})
	}
}
