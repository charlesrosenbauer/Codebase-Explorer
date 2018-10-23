package main



import (
  //"math"
)




func (c *ClusterData) GetFitness() float64 {
    // Let's start out with something simple for now. If there are any nodes
    // outside the window, or closer than a specific distance from each other,
    // decrease the fitness.

    fitness := 1.0
    for i, f := range c.Files {

      if f.X < -0.9 || f.X > 0.9 || f.Y < -0.9 || f.Y > 0.9 {
        fitness *= 0.95
      }

      for j, g := range c.Files {
        if i < j {
          dx := (f.X - g.X)
          dy := (f.Y - g.Y)

          if (dx * dx) + (dy * dy) < 0.005 {
            fitness *= 0.95
          }
        }
      }
    }
    return fitness
}
