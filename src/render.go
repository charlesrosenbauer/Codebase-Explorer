package main



import (
  "math/rand"
  "github.com/veandco/go-sdl2/sdl"
)





const HEIGHT = 600
const WIDTH  = 800







func DrawNode(x, y int32, highlight uint8, s *sdl.Surface){
  r := sdl.Rect{x, y, 5, 5}
  s.FillRect(&r, uint32(255 - highlight))
}







func (c *ClusterData)Randomize(){
  for _, v := range c.Files {
    v.X = rand.Float32()
    v.Y = rand.Float32()
  }
}






func (c *ClusterData)MoveCluster(){
  for i, _ := range c.Files {
    for j, _ := range c.Edges[i] {
      x := c.Files[j].X
      y := c.Files[j].Y
      c.Files[i].X += x * c.Edges[i][j]
      c.Files[i].Y += y * c.Edges[i][j]
    }
  }
}





func (c *ClusterData)DrawCluster(s *sdl.Surface){
  s.FillRect(nil, 0)
  for _, v := range c.Files {
    DrawNode(int32(v.X * WIDTH), int32(v.Y * HEIGHT), 255, s)
  }
}
