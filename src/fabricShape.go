package main

import (
	"strings"

	"github.com/fogleman/gg"
)

type fabricShape struct {
	ShapeType   string  `json:"type"`
	Left        float64 `json:"left"`
	Top         float64 `json:"top"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Fill        string  `json:"fill"`
	Stroke      string  `json:"stroke"`
	StrokeWidth float64 `json:"strokeWidth"`
	Angle       float64 `json:"angle"`
	ScaleX      float64 `json:"scaleX"`
	ScaleY      float64 `json:"scaleY"`
	Radius      float64 `json:"radius"`
	Text        string  `json:"text"`
	X1          float64 `json:"x1"`
	X2          float64 `json:"x2"`
	Y1          float64 `json:"y1"`
	Y2          float64 `json:"y2"`
}

const (
	circle    = "circle"
	rectangle = "rect"
	line      = "line"
	text      = "i-text"
)

func (s *fabricShape) drawShapeType(ctx *gg.Context) {
	switch s.ShapeType {
	case circle:
		ctx.DrawCircle(s.Left, s.Top, s.Radius)
		break

	case rectangle:
		ctx.DrawRectangle(s.Left, s.Top, s.Width, s.Height)
		break

	case line:
		ctx.DrawLine(s.X1, s.Y1, s.X2, s.Y2)
		break

	case text:
		ctx.DrawString(s.Text, s.Left, s.Top)
		break
	}
}

// Adds the current shape to the gg canvas
func (s *fabricShape) Parse(ctx *gg.Context) {
	ctx.Push()
	if s.ScaleX != 1 || s.ScaleY != 1 {
		ctx.ScaleAbout(s.ScaleX, s.ScaleY, s.Left, s.Top)
	}

	if s.Angle != 0 {
		ctx.RotateAbout(gg.Radians(s.Angle), s.Left, s.Top)
	}

	s.drawShapeType(ctx)

	if len(s.Stroke) > 0 && strings.ToLower(s.Stroke) != "transparent" {
		ctx.SetHexColor(s.Stroke)
		ctx.SetLineWidth(s.StrokeWidth)

		ctx.StrokePreserve()
	}

	if len(s.Fill) > 0 && strings.ToLower(s.Fill) != "transparent" {
		ctx.SetHexColor(s.Fill)
		ctx.Fill()
	}

	ctx.Pop()
}