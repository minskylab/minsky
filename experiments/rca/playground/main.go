package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"time"

	"github.com/minskylab/rca"
	"github.com/minskylab/rca/gol"
)

func main() {
	images := make(chan image.Image)

	model := gol.NewGoLDynamicalSystem(128, 128, time.Now().Unix())

	renderer := gol.NewImageRenderer(images)

	vm := rca.NewVM(model, renderer)

	go func() {
		i := 0
		for image := range images {
			fmt.Println(image.Bounds())

			f, err := os.OpenFile("img_"+strconv.Itoa(i)+".jpeg", os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}

			if err := jpeg.Encode(f, image, nil); err != nil {
				panic(err)
			}

			f.Close()
		}

	}()

	done := vm.Run(1 * time.Millisecond)

	<-done
}
