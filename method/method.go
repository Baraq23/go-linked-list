package method1


type Rectangle struct {
	Width float64
	Length float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.Length + r.Width)
}

func (r Rectangle) Volume(h float64) float64 {
	vol := h * r.Length *r.Width
	return vol
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}