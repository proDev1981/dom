package main

import "app/dom"
import "app/ui"

func main(){
  dom.New(
    dom.RenderDom(ui.App()),
    dom.NewWindow().
      SetTitle("New app"),
    )
  dom.OnWait()
}
