package main

import (
	"log"
	"os"

	"github.com/stevenferrer/nero/gen"

	"github.com/stevenferrer/nero-example/model"
)

func main() {
	files, err := gen.Generate((model.Product{}).Schema())
	if err != nil {
		log.Fatal(err)
	}

	basePath := "productrepo"
	err = os.Mkdir(basePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = file.Render(basePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
