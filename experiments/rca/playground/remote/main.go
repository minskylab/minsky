package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"time"

	"github.com/minskylab/rca"
	"github.com/minskylab/rca/cyclic"
	"github.com/minskylab/rca/remote"
)

func main() {
	images := make(chan image.Image)

	width, height := 512, 512

	// model := gol.NewGoLDynamicalSystem(width, height, time.Now().Unix())

	// renderer := gol.NewImageRenderer(images, width, height)

	model, err := cyclic.NewRockPaperSissor(width, height, 2, time.Now().Unix(), images)
	if err != nil {
		panic(err)
	}

	vm := rca.NewVM(rca.BulkDynamicalSystem(model, model), model)

	dataSource := make(chan []byte)

	go func(dataSource chan []byte) {
		buff := bytes.NewBuffer([]byte{})

		for img := range images {
			if err := jpeg.Encode(buff, img, &jpeg.Options{
				Quality: 100,
			}); err != nil {
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
