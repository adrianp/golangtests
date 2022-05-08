package utils

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
)

func Lissajous(out io.Writer) {
    const (
        whiteIndex = 0
        blackIndex = 1
        cycles = 5
        angularResolution = 0.001
        size = 100
        frames = 64
        delay = 8
    )

    var palette = []color.Color{color.White, color.Black}

    frequency := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: frames}
    phase := 0.0
    for i := 0; i < frames; i++ {
        rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rectangle, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += angularResolution {
            x := math.Sin(t)
            y := math.Sin(t*frequency + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}

