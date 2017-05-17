// Package bytes は、KB, MB, ..., YB までの定数を提供します。
package bytes

const (
	// KB は、キロバイトを表します。
	KB = 1000
	// MB は、メガバイトを表します。
	MB = KB * KB
	// GB は、ギガバイトを表します。
	GB = MB * KB
	// TB は、テラバイトを表します。
	TB = GB * KB
	// PB は、ペタバイトを表します。
	PB = TB * KB
	// EB は、エクサバイトを表します。
	EB = PB * KB
	// ZB は、ゼタバイトを表します。
	ZB = EB * KB
	// YB は、ヨタバイトを表します。
	YB = ZB * KB
)
