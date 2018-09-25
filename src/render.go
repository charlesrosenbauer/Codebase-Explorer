package main



import (
  "github.com/veandco/go-sdl2/sdl"
)






func (s *Surface)DrawNode(x, y int, highlight uint8){
  r := Rect{x, y, 5, 5}
  s.FillRect(r, uint32(255 - highlight))
}
