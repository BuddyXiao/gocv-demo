/*
 * @Author: 肖和锋
 * @Date: 2022/7/7 15:34
 */
package common

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"testing"
)

func TestOne(t *testing.T) {
	src := gocv.IMRead("baihe.jpg", gocv.IMReadColor)
	mean := [3]float32{0.485, 0.456, 0.406}
	std := [3]float32{0.229, 0.224, 0.225}
	Normalize(src,mean,std)
}

func TestTwo(t *testing.T) {
	src := gocv.IMRead("baihe.jpg", gocv.IMReadColor)
	mean := [3]float32{0.485, 0.456, 0.406}
	std := [3]float32{0.229, 0.224, 0.225}
	gocv.Resize(src,&src,image.Point{224,224},0,0,gocv.InterpolationLinear)
	blobFromImage := Normalize(src,mean,std)
	blobFromImage.ConvertTo(blobFromImage,gocv.MatTypeCV32FC3)
	var w = blobFromImage.Cols()
	var h = blobFromImage.Rows()
	var c = blobFromImage.Channels()
	fmt.Println(w,h,c)
	traverse(*blobFromImage)
}

func TestMinMax(t *testing.T) {
	src := gocv.IMRead("baihe.jpg", gocv.IMReadColor)
	split := gocv.Split(src)
	idx, maxVal, minIdx, i := gocv.MinMaxIdx(split[0])
	println(idx,maxVal,minIdx,i)
}