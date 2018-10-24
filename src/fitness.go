package main



import (
  //"math"
  "math/rand"
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




// Makes some random small changes to the cluster model. If the fitness is better
// than the provided value, it updates the cluster model with the changes.
// Returns a new version of ClusterData, and whether or not the step was made.
func TryStep(c *ClusterData, fitnessTarget, stepSize float64) (*ClusterData, bool) {
  ret := *c
  for _, f := range ret.Files {
    f.X += float32(stepSize) * (0.5 - rand.Float32())
    f.Y += float32(stepSize) * (0.5 - rand.Float32())
  }

  fitness := ret.GetFitness()
  if fitness > fitnessTarget {
    return &ret, true
  }

  return c, false
}
