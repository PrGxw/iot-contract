package contract

import (
	"fmt"
	"strings"
)

// NormalizeMAC returns the canonical MAC form: uppercase, colon-separated
// ("AA:BB:CC:DD:EE:FF"), the form the backend stores. It accepts 12 hex digits
// with or without ':' or '-' separators, in any case.
func NormalizeMAC(s string) (string, error) {
	compact := strings.NewReplacer(":", "", "-", "").Replace(s)
	if len(compact) != 12 {
		return "", fmt.Errorf("invalid MAC %q", s)
	}
	for _, r := range compact {
		if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F') {
			return "", fmt.Errorf("invalid MAC %q", s)
		}
	}
	compact = strings.ToUpper(compact)
	var b strings.Builder
	for i := 0; i < 12; i += 2 {
		if i > 0 {
			b.WriteByte(':')
		}
		b.WriteString(compact[i : i+2])
	}
	return b.String(), nil
}
