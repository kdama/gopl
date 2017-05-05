// Package bytes は、KB, MB, ..., YB までの定数を提供します。
package bytes

const (
	// KiB は、キビバイトを表します。
	KiB = 1 << (10 * iota)
	// MiB は、メビバイトを表します。
	MiB
	// GiB は、ギビバイトを表します。
	GiB
	// TiB は、テビバイトを表します。
	TiB
	// PiB は、ペビバイトを表します。
	PiB
	// EiB は、エクスビバイトを表します。
	EiB
	// ZiB は、ゼビバイトを表します。
	ZiB
	// YiB は、ヨビバイトを表します。
	YiB
)
