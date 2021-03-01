package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"time"

	"github.com/minskylab/rca"
	"github.com/minskylab/rca/gol"
	"github.com/minskylab/rca/remote"
)

func main() {
	images := make(chan image.Image)

	width, height := 512, 512

	model := gol.NewGoLDynamicalSystem(width, height, time.Now().Unix())

	renderer := gol.NewImageRenderer(images, width, height)

	vm := rca.NewVM(model, renderer)

	dataSource := make(chan []byte)

	go func(dataSource chan []byte) {
		buff := bytes.NewBuffer([]byte{})

		for img := range images {
			// fmt.Printf("enconding %p\n", img)
			if err := jpeg.Encode(buff, img, nil); err != nil {
				panic(err)
			}
			dataSource <- buff.Bytes()
			buff.Reset()
		}
	}(dataSource)

	rs := remote.NewBinaryRemote(3000, "/", dataSource)

	go rs.Run()

	vm.Run(1000 * time.Second)
}
