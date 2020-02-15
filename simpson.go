// シンプソン法

package main

func simpson(a, b float64, n int64, f func(float64) float64) float64 {
	delta := (b - a) / float64(n)
	s := float64(0.0)
	for x := a; x < b; x += delta {
		s += _simpson(x, x+delta, f)
	}
	return s
}

func _simpson(a, b float64, f func(float64) float64) float64 {
	return (b - a) * (f(a) + float64(4.0)*f((a+b)/float64(2.0)) + f(b)) / float64(6.0)
}
