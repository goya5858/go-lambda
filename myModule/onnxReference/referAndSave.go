package onnxReference

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"myModule/images"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
)

func OnnxRef(inpath string, outpath string) {
	// I. Inputの作成
	file, fileOpenErr := os.Open(inpath)
	defer file.Close()
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	}

	img, format, decodeErr := image.Decode(file)
	fmt.Println(format) // PNG
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	height, width := img.Bounds().Dy(), img.Bounds().Dx()
	input := tensor.New(tensor.WithShape(1, 3, height, width), tensor.Of(tensor.Float32))
	convertErr := images.ImageToBCHW(img, input)
	if convertErr != nil {
		fmt.Println("ImageToBCHW error:", convertErr)
	}
	//fmt.Println(input)

	// II. Modelの作成
	backend := gorgonnx.NewGraph()
	model := onnx.NewModel(backend)

	byte_model, _ := os.ReadFile("../SampleModel.onnx")
	ReadModelErr := model.UnmarshalBinary(byte_model)
	if ReadModelErr != nil {
		fmt.Println(ReadModelErr)
	}

	// III. Inference
	model.SetInput(0, input)
	runErr := backend.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
	output, _ := model.GetOutputTensors()

	outimg, convertErr2 := images.TensorToImg(output[0])
	if convertErr2 != nil {
		log.Fatal(convertErr2)
	}
	//fmt.Println("outIMG:", outimg)

	// IV. 予測の保存
	file_out, createErr := os.Create(outpath)
	defer file_out.Close()
	if createErr != nil {
		log.Fatal(createErr)
	}
	if format == "png" {
		png.Encode(file_out, outimg)
	} else {
		jpeg.Encode(file_out, outimg, nil)
	}

}
