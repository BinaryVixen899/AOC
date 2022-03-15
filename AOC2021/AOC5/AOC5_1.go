package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()
	p.Title.Text = "X"
	p.Title.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", addPoints(0, 9, 5, 9),
		"Second", addPoints(8, 0, 0, 8),
		// These need to be done programatically
		"Third", addPoints(9, 4, 3, 4),
		"Fourth", addPoints(2, 2, 2, 1),
		"Fifth", addPoints(7, 0, 7, 4),
		"Sixth", addPoints(6, 4, 2, 0),
		"Seventh", addPoints(0, 9, 2, 9),
		"Eight", addPoints(0, 0, 8, 8),
		"Ninth", addPoints(5, 5, 8, 2))
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

}

func addPoints(sx float64, sy float64, fx float64, fy float64) plotter.XYs {
	pts := make(plotter.XYs, 4)
	for i := range pts {
		if i == 0 || i == 1 {
			pts[i].X = sx
			pts[i].Y = sy
		} else {
			pts[i].X = fx
			pts[i].Y = fy
		}

	}
	return pts

}
