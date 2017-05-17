package bytes

import (
	"testing"
)

func TestConsts(t *testing.T) {
	if got := KB; got != 1000 {
		t.Errorf("KB = %d, want 1000", got)
	}
	if got := MB; got != 1000000 {
		t.Errorf("MB = %d, want 1000000", got)
	}
	if got := MB / KB; got != 1000 {
		t.Errorf("MB / KB = %d, want 1000", got)
	}
	if got := TB / MB; got != 1000000 {
		t.Errorf("TB / MB = %d, want 1000000", got)
	}
}
