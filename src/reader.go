package main


import(
  "regexp"
  "math"
  "strings"
  "math/rand"
)



type ProgramMap struct{
  Filepath string
  Wordcts  map[string]int
  Wordwts  map[string]float32
  Totalwt  float32
  Totalct  int
}



func MapFile(text string, path string) ProgramMap {
  var ret ProgramMap
  ret.Filepath = path
  ret.Totalct = 0

  regexpr := "[a-zA-Z\\d_]+"
  regex, err := regexp.Compile(regexpr)
  if err != nil {
    panic(err)
  }

  textlines := strings.Split(text, "\n")
  for i:=0; i < len(textlines); i++ {
    ids := regex.FindAllString(textlines[i], -1)
    for _, v := range ids {
      if _, ok := ret.Wordcts[v]; ok {
        ret.Wordcts[v] ++
      }else{
        ret.Wordcts[v] = 1
      }
      ret.Totalct++
    }
  }
  return ret
}




func (p *ProgramMap)GetLocalWeights(){
  factor := float32(math.Log(float64(p.Totalct)))
  factor *= (factor / float32(p.Totalct))
  for i, v := range p.Wordcts {
    val := float32(v)

    val = val * factor
    p.Wordwts[i] = val
    p.Totalwt += val
  }
}



func MakeGlobalMap(maps []ProgramMap) ProgramMap {
  var ret ProgramMap
  ret.Filepath = ""
  ret.Totalct = 0

  for _, m := range maps {
    for v, _ := range m.Wordcts {
      if _, ok := ret.Wordcts[v]; ok {
        ret.Wordcts[v]++
      }else{
        ret.Wordcts[v] = 1
      }
      ret.Totalct += m.Wordcts[v]
    }
  }
  return ret
}


func ComputeCorrelation(a, b *ProgramMap) float32 {

  var x map[string]float32
  var y map[string]float32
  var totalA float32
  var totalB float32
  if len(a.Wordwts) < len(b.Wordwts) {
    x = a.Wordwts
    y = b.Wordwts
    totalA = a.Totalwt
    totalB = b.Totalwt
  }else{
    x = b.Wordwts
    y = a.Wordwts
    totalA = b.Totalwt
    totalB = a.Totalwt
  }

  var accuma float32 = 0.0
  var accumb float32 = 0.0
  for i, v := range x {
    if wgt, ok := y[i]; ok {
      accuma += wgt
      accumb += v
    }
  }

  parta := accuma / totalA
  partb := accumb / totalB

  return float32(math.Sqrt(float64(parta * partb)))
}





type FileData struct{
  Path string
  X, Y float32
}

type ClusterData struct{
  Files []FileData
  Edges [][]float32
  Size  int
}



func MakeClusterData(maps []ProgramMap) ClusterData {

  var ret ClusterData
  ret.Size = len(maps)

  ret.Files = make([]FileData, ret.Size)
  ret.Edges = make([][]float32, ret.Size)
  for i := 0; i < ret.Size; i++ {
    ret.Edges[i] = make([]float32, ret.Size);
  }

  for i, va := range maps {
    for j, vb := range maps {
      if i < j {
        ret.Edges[i][j] = ComputeCorrelation(&va, &vb)
      }else{
        ret.Edges[i][j] = 0.0
      }
    }
    ret.Files[i] = FileData{va.Filepath, rand.Float32(), rand.Float32()}
  }
  return ret
}
