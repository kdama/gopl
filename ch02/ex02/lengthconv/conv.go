// Package lengthconv は、フィートとメートルの長さの計算を行います。
package lengthconv

// FootToMeter は、フィートの長さをメートルの長さに変換します。
func FootToMeter(f Foot) Meter { return Meter(float64(f) / FootPerMeter) }

// MeterToFoot は、メートルの長さをフィートの長さに変換します。
func MeterToFoot(m Meter) Foot { return Foot(float64(m) * FootPerMeter) }
