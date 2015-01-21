package main



import (
    "9fans.net/go/draw"
    "image"
    "fmt"
)

func main() {
  d, err := draw.Init(nil, "", "gnal", "512x512")
  if err != nil { panic(err) }
  
  screen := d.ScreenImage
  var pstring image.Point
  var s string

  for i := 0; ; i++ {
    s = fmt.Sprintf("%d", i)
    screen.Draw(image.Rectangle{pstring, pstring.Add(d.DefaultFont.StringSize(s))}, d.White, nil, image.ZP)
    screen.String(pstring, d.Black, image.ZP, d.DefaultFont, s)
    d.Flush()
  }

}
