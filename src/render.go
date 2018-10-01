package main



import (
  "math"
  "math/rand"
  "github.com/veandco/go-sdl2/sdl"
  "fmt"
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
  rngx := 1.0
  rngy := 1.0
  for i, _ := range c.Files {
    for j, _ := range c.Edges[i] {
      if i < j {
        x := c.Files[j].X
        y := c.Files[j].Y
        c.Files[i].X += x * c.Edges[i][j] * t
        c.Files[i].Y += y * c.Edges[i][j] * t
      }
    }
    if math.Abs(float64(c.Files[i].X)) > rngx {
      rngx = float64(c.Files[i].X)
    }
    if math.Abs(float64(c.Files[i].Y)) > rngy {
      rngy = float64(c.Files[i].Y)
    }
  }

  for _, f := range c.Files {
    f.X /= float32(rngx)
    f.Y /= float32(rngy)
  }
}





func (c *ClusterData)DrawCluster(s *sdl.Surface){
  s.FillRect(nil, 0)
  for _, v := range c.Files {
    DrawNode(int32((v.X + 1.0) * (WIDTH / 2)), int32((v.Y + 1.0) * (HEIGHT / 2)), 0, s)
    fmt.Println(v.X, v.Y, v.Path)
  }
  fmt.Println("----")
}
