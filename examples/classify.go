/*
 * @Author: 肖和锋
 * @Date: 2022/7/7 14:41
 */
package examples

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"log"
)

const (
	w = 224
	h
)

func classify(imgPath string) string{
	//imgPath := "config/baihe.jpg"
	modelPath := "config/model.onnx"
	imgSrc := gocv.IMRead(imgPath, gocv.IMReadColor)
	// 图片预处理
	processedImg := preprocess(&imgSrc)
	onnxNet := gocv.ReadNetFromONNX(modelPath)
	if onnxNet.Empty() {
		log.Fatal("读取网络模型失败...")
	}
	onnxNet.SetInput(processedImg,"image")
	result := onnxNet.Forward("softmax_0.tmp_0")
	return performClassify(result)
}


func performClassify(result gocv.Mat) string{
	classes := []string {"baihe", "dangshen","gouqi","huaihua","jinyinhua"}
	ptrFloat32, err := result.DataPtrFloat32()
	if err != nil {
		return ""
	}
	maxIdx,ratio := argMax(ptrFloat32)
	fmt.Println(ptrFloat32)
	fmt.Printf("class: %s, ratio: %f\n",classes[maxIdx],ratio)
	return classes[maxIdx]
}

func argMax(slice []float32) (int,float32) {
	var maxIdx = 0
	var ratio = slice[maxIdx]
	for i := 0; i < len(slice); i++ {
		if slice[i] > slice[maxIdx]{
			maxIdx = i
			ratio = slice[i]
		}
	}
	return maxIdx,ratio
}

func preprocess(src *gocv.Mat) gocv.Mat{
	if src.Empty() {
		log.Fatal("图片有误...")
	}
	blobFromImage := gocv.BlobFromImage(*src, 1/255.0, image.Point{
		X: w,
		Y: h,
	}, gocv.Scalar{}, true, false)
	return blobFromImage
}

