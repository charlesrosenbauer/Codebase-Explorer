package main



import (
  "github.com/veandco/go-sdl2/sdl"
)







func DrawNode(x, y int32, highlight uint8, s *sdl.Surface){
  r := sdl.Rect{x, y, 5, 5}
  s.FillRect(&r, uint32(255 - highlight))
}
