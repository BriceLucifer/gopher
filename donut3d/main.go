package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	A, B := 0.0, 0.0
	var i, j float64
	var k int
	z := make([]float64, 1760)
	b := make([]byte, 1760)
	fmt.Print("\x1b[2J")
	for {
		// 初始化 b 和 z
		for i := range b {
			b[i] = ' '
		}
		for i := range z {
			z[i] = 0
		}

		for j = 0; j < 6.28; j += 0.07 {
			for i = 0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e
				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := x + 80*y
				N := int(8 * ((f*e - c*d*g) * m - c*d*e - f*g - l*d*n))
				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 && N < len(",./@#&^*$#") {
						b[o] = ",./@#&^*$#"[N]
					} else {
						b[o] = '$'
					}
				}
			}
		}

		fmt.Print("\x1b[H")
		for k = 0; k < 1760; k++ {
			if k%80 == 0 {
				fmt.Print("\n")
			} else {
				fmt.Printf("%c", b[k])
			}
		}

		A += 0.04  // 增加旋转速度
		B += 0.02 // 增加旋转速度
		time.Sleep(8 * time.Millisecond) // 减少睡眠时间，提高帧率
	}
}
