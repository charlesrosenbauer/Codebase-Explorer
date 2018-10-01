package main



import (
  //"math"
  "math/rand"
  "github.com/veandco/go-sdl2/sdl"
  //"fmt"
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






func (c *ClusterData)MoveCluster(t float32){
  var minX float32 =  100.0
  var maxX float32 = -100.0
  var minY float32 =  100.0
  var maxY float32 = -100.0
  for i, _ := range c.Files {
    for j, _ := range c.Edges[i] {
      if i < j {
        x := c.Files[j].X
        y := c.Files[j].Y
        dx := x - c.Files[i].X
        dy := y - c.Files[i].Y
        if dx < 0.05 {
          dx *= -0.3
        }
        if dy < 0.05 {
          dy *= -0.3
        }
        c.Files[i].X += dx * c.Edges[i][j] * t
        c.Files[i].Y += dy * c.Edges[i][j] * t
      }
    }
    if c.Files[i].X > maxX {
      maxX = c.Files[i].X
    }else if c.Files[i].X < minX {
      minX = c.Files[i].X
    }
    if c.Files[i].Y > maxY {
      maxY = c.Files[i].Y
    }else if c.Files[i].Y < minY {
      minY = c.Files[i].Y
    }
  }

  midX := (maxX + minX) / 2
  midY := (maxY + minY) / 2
  rngX :=  maxX - minX
  rngY :=  maxY - minY

  for i, _ := range c.Files {
    c.Files[i].X = (c.Files[i].X - midX) / rngX
    c.Files[i].Y = (c.Files[i].Y - midY) / rngY
  }
}





func (c *ClusterData)DrawCluster(s *sdl.Surface){
  s.FillRect(nil, 0)
  for _, v := range c.Files {
    DrawNode(int32((v.X + 1.0) * (WIDTH / 2)), int32((v.Y + 1.0) * (HEIGHT / 2)), 0, s)
    //fmt.Println(v.X, v.Y, v.Path)
  }
  //fmt.Println("----")
}
