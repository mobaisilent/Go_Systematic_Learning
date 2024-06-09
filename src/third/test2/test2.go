package main

import "fmt"

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1), which is 1024
	MB                             // 1 << (10*2), which is 1048576
	GB                             // 1 << (10*3), which is 1073741824
	TB                             // 1 << (10*4), which is 1099511627776
	PB                             // 1 << (10*5), which is 1125899906842624
	EB                             // 1 << (10*6), which is 1152921504606846976
	ZB                             // 1 << (10*7), which is too large for a float64
	YB                             // 1 << (10*8), which is too large for a float64
)

func main() {
	fmt.Printf("KB: %f\n", KB)
	fmt.Printf("MB: %f\n", MB)
	fmt.Printf("GB: %f\n", GB)
	fmt.Printf("TB: %f\n", TB)
	fmt.Printf("PB: %f\n", PB)
	fmt.Printf("EB: %f\n", EB)
	// Note: ZB and YB values are too large to represent accurately with float64
	// and will result in an overflow, leading to incorrect values when printed.
	// Hence, they are not printed here.
}

/*
打印的结果为：
KB: 1024.000000
MB: 1048576.000000
GB: 1073741824.000000
TB: 1099511627776.000000
PB: 1125899906842624.000000
EB: 1152921504606846976.000000
*/
