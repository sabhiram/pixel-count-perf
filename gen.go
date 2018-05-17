package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math/rand"
	"os"
)

func randomImage(w, h int, percent float64) (*image.RGBA, error) {
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	im := image.NewRGBA(r)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if rand.Float64() <= percent {
				im.Set(i, j, color.RGBA{
					R: byte(rand.Intn(255)),
					G: byte(rand.Intn(255)),
					B: byte(rand.Intn(255)),
					A: 255,
				})
			} else {
				im.Set(i, j, color.White)
			}
		}
	}
	return im, nil
}

func randomMonochromeImage(w, h int, percent float64) (*image.RGBA, error) {
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	im := image.NewRGBA(r)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if rand.Float64() <= percent {
				v := byte(rand.Intn(255))
				im.Set(i, j, color.RGBA{
					R: v,
					G: v,
					B: v,
					A: 255,
				})
			} else {
				im.Set(i, j, color.White)
			}
		}
	}
	return im, nil
}

func fatalOnError(err error) {
	if err != nil {
		fmt.Printf("Fatal error :: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	fatalOnError(os.MkdirAll("./gen", 0655))

	for _, tc := range []struct {
		w, h int
		p    float64
		fn   string
	}{
		{128, 128, 0.00, "gen/small_00p_0"},
		{128, 128, 0.25, "gen/small_25p_0"},
		{128, 128, 0.50, "gen/small_50p_0"},
		{128, 128, 0.75, "gen/small_75p_0"},
		{128, 128, 0.95, "gen/small_95p_0"},
		{128, 128, 1.00, "gen/small_100p_0"},

		{2048, 2048, 0.00, "gen/large_00p_0"},
		{2048, 2048, 0.25, "gen/large_25p_0"},
		{2048, 2048, 0.50, "gen/large_50p_0"},
		{2048, 2048, 0.75, "gen/large_75p_0"},
		{2048, 2048, 0.95, "gen/large_95p_0"},
		{2048, 2048, 1.00, "gen/large_100p_0"},

		{2048 * 4, 2048 * 4, 0.00, "gen/behemoth_00p_0"},
		{2048 * 4, 2048 * 4, 0.25, "gen/behemoth_25p_0"},
		{2048 * 4, 2048 * 4, 0.50, "gen/behemoth_50p_0"},
		{2048 * 4, 2048 * 4, 0.75, "gen/behemoth_75p_0"},
		{2048 * 4, 2048 * 4, 0.95, "gen/behemoth_95p_0"},
		{2048 * 4, 2048 * 4, 1.00, "gen/behemoth_100p_0"},
	} {
		for _, desc := range []struct {
			fn  func(int, int, float64) (*image.RGBA, error)
			pfx string
		}{
			{randomImage, "color"},
			{randomMonochromeImage, "bw"},
		} {
			im, err := desc.fn(tc.w, tc.h, tc.p)

			// pngFn := fmt.Sprintf("%s_%s.png", tc.fn, desc.pfx)
			// f, err := os.Create(pngFn)
			// fatalOnError(err)
			// defer f.Close()
			// fatalOnError(png.Encode(f, im))
			// fmt.Printf("Wrote file %s\n", pngFn)

			rawFn := fmt.Sprintf("%s_%s.raw", tc.fn, desc.pfx)
			ioutil.WriteFile(rawFn, im.Pix, 0655)
			fmt.Printf("Wrote file %s\n", rawFn)
		}
	}
}
