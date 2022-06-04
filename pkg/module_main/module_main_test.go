package module_main

import "testing"

func TestDoingSomeInteresting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		struct {
			name string
			want string
		}{"ok", "1.00001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoingSomeInteresting(); got != tt.want {
				t.Errorf("DoingSomeInteresting() = %v, want %v", got, tt.want)
			}
		})
	}
}
