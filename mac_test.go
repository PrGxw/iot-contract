package contract

import "testing"

func TestNormalizeMAC_Forms(t *testing.T) {
	for _, in := range []string{"aa:bb:cc:dd:ee:ff", "AABBCCDDEEFF", "aa-bb-cc-dd-ee-ff", "Aa:bB:cC:dD:eE:fF", "aabbccddeeff"} {
		got, err := NormalizeMAC(in)
		if err != nil || got != "AA:BB:CC:DD:EE:FF" {
			t.Errorf("NormalizeMAC(%q) = %q, %v", in, got, err)
		}
	}
}

func TestNormalizeMAC_Invalid(t *testing.T) {
	for _, in := range []string{"zz", "", "AA:BB:CC:DD:EE", "AA:BB:CC:DD:EE:FF:00", "GG:BB:CC:DD:EE:FF", "AABBCCDDEEF"} {
		if got, err := NormalizeMAC(in); err == nil {
			t.Errorf("NormalizeMAC(%q) = %q, want error", in, got)
		}
	}
}
