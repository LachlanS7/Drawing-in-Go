package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
	"time"
)

func main() {
	start := time.Now()

	makeImage(1024, 1024, -0.8, 0.8, -0.8, 0.8, "Hello")

	end := time.Now()

	fmt.Println(end.Sub(start).Milliseconds())
}

func makeImage(width, height int, lowerDomain, upperDomain, lowerRange, upperRange float64, fileName string) {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	spanX := upperDomain - lowerDomain
	spanY := upperRange - lowerRange

	for y := 0; y < height; y++ {

		for x := 0; x < width; x++ {
			img.Set(x, y, mandelbrot(lowerDomain+float64(spanX)*(float64(x)/float64(width)), upperRange-spanY*(float64(y)/float64(height))))
		}
	}

	f, err := os.Create(fileName + ".bmp")

	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func mandelbrot(x, y float64) color.NRGBA {

	temp := 0 + 0i

	tempPos := 0
	tempValue := 0 + 0i

	for i := 0; i < 100; i++ {

		temp = cmplx.Pow(temp, 2) //cmplx.Pow(cmplx.Conj(temp), complex(1/x, -1/y))

		//temp *= cmplx.Log(complex(y, math.Pow(x, 1/2)))

		temp += complex(x, y)

		if i == 50 {
			tempPos = 50
			tempValue = temp
		}

		if i > 50 {
			if math.Abs(real(temp)-real(tempValue)) < 0.05 && math.Abs(imag(temp)-imag(tempValue)) < 0.05 {
				return color.NRGBA{255 - uint8(40*(i-tempPos)), uint8(40 * (i - tempPos)), 0, 255}
			}
		}

		if cmplx.Abs(temp) > 2 {
			return color.NRGBA{0, 0, 0, 255}
		}
	}
	return color.NRGBA{255, 0, 0, 255}
}
