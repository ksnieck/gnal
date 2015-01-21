package main



import (
    "9fans.net/go/draw"
    "image"
)

func main() {
  d, err := draw.Init(nil, "", "gnal", "512x512")
  if err != nil { panic(err) }
  
  screen := d.ScreenImage


  stringz := d.DefaultFont.StringSize("000000")
  var pstring image.Point

  for {
    screen.Draw(image.Rectangle{pstring, pstring.Add(stringz)}, d.White, nil, image.ZP)
    screen.String(pstring, d.Black, image.ZP, d.DefaultFont, "gnal")
    d.Flush()
  }




}
