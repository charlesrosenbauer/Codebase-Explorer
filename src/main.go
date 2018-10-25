package main


import (
  "fmt"
  "os"
  "io/ioutil"
  "github.com/veandco/go-sdl2/sdl"
)



func loadSourceFiles(files []string) ([]string, error) {
	var filetexts []string
	for _, file := range files {
		text, err := ioutil.ReadFile(file)
		if err != nil {
			out := []string{}
			return out, err
		}
		filetexts = append(filetexts, string(text))
	}
	return filetexts, nil
}






func main(){

  files := os.Args[1:]
  text, err := loadSourceFiles(files)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%d files loaded.\n", len(text))

  if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
    panic(err)
  }
  defer sdl.Quit()

	window, err := sdl.CreateWindow("Codebase Explorer", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

  maps:= make([]ProgramMap, len(text))
  for i, txt := range text {
    maps[i] = MapFile(txt, files[i])
  }

  //gmap := MakeGlobalMap(maps)

  cdata := MakeClusterData(maps)

  for i := 0; i < 5000; i++ {
    window.UpdateSurface()
    cdata.DrawCluster(surface)
    //cdata.MoveCluster(0.005)
    sdl.Delay(16)
  }
}
