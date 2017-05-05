package bytes

import (
	"testing"
)

func TestConsts(t *testing.T) {
	if got := MiB / KiB; got != 1024 {
		t.Errorf("GiB / KiB = %d, want 1024", got)
	}
	if got := TiB / MiB; got != 1048576 {
		t.Errorf("TiB / MiB = %d, want 1048576", got)
	}
}
