package constants

type PaperSize struct {
	Width         float64
	Height        float64
	WidthInPixel  float64
	HeightInPixel float64
	Name          string
	Dpi           float64
}

var A4 PaperSize = PaperSize{Width: 210, Height: 297, WidthInPixel: 2480, HeightInPixel: 3508, Name: "A4", Dpi: 300}
var A5 PaperSize = PaperSize{Width: 148, Height: 210, WidthInPixel: 1748, HeightInPixel: 2480, Name: "A5", Dpi: 300}
