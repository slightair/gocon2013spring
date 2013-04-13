package main

import (
  "os"
  "fmt"
  "image"
  "image/color"
  "image/png"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("usage: go run color_changer.go <infile> <outfile>")
    os.Exit(1)
  }
  
  in  := os.Args[1]
  out := os.Args[2]
  
  infile, err := os.Open(in)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  base, err := png.Decode(infile)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  rect := base.Bounds()
  
  palette := color.Palette{}
  for y := 0; y < rect.Dy(); y++ {
    for x := 0; x < rect.Dx(); x++ {
      c := base.At(x, y)
      
      isUnique := true
      for _, e := range palette {
        if e == c {
          isUnique = false
          break
        }
      }
      
      if isUnique {
        palette = append(palette, c)
      }
    }
  }
  
  source := image.NewPaletted(rect, palette);
  for y := 0; y < rect.Dy(); y++ {
    for x := 0; x < rect.Dx(); x++ {
      c := base.At(x, y)
      source.Set(x, y, c)
    }
  }
  
  palette[0] = color.RGBA{255, 0, 255, 255}
  dest := image.Paletted{source.Pix, source.Stride, source.Rect, palette}
  
  outfile, err := os.OpenFile(out, os.O_WRONLY | os.O_CREATE, 0644)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  err = png.Encode(outfile, &dest)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}