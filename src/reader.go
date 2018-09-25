package main


import(
  "regexp"
  "math"
  "strings"
)



type ProgramMap struct{
  Filepath string
  Wordcts  map[string]int
  Wordwts  map[string]float32
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
