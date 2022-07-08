/*
 * @Author: 肖和锋
 * @Date: 2022/7/7 15:14
 */
package common

import (
	"fmt"
	"gocv.io/x/gocv"
)

// Normalize 读取的图片是 rgb通道的图片 ,
// im = (im - min_value) / (max_value - min_value)
// x' =  ( x - mean ) / std

func Normalize(mat gocv.Mat, mean [3]float32, std [3]float32) *gocv.Mat {
	// 将读取的图片转换为 rgb
	mat_rgb := mat.Clone()
	BGR2RGB(mat,&mat_rgb)
	dst := gocv.NewMatWithSize(mat_rgb.Rows(),mat_rgb.Cols(),gocv.MatTypeCV32FC3) // 默认是uchar类型
	norm_zscore(mat_rgb,&dst,mean,std)
	var w = dst.Cols()
	var h = dst.Rows()
	var c = dst.Channels()
	fmt.Println(w,h,c)
	//traverse(dst)
	return &dst
}

func norm_zscore(src gocv.Mat, dst *gocv.Mat,mean [3]float32, std [3]float32) {
	var w = dst.Cols()
	var h = dst.Rows()
	var c = dst.Channels()
	split := gocv.Split(src)
	for ci := 0; ci < c; ci++ {
		min, max, _, _ := gocv.MinMaxIdx(split[ci])
		for x := 0; x < h; x++ {
			for y := 0; y < w; y++ {
				p := src.GetUCharAt3(x,y,ci)
				tmp := float32((float32(p)-min)/(max-min))
				r := float32(( tmp - mean[ci] ) / std[ci])
				dst.SetFloatAt3(x,y,ci,r)
			}
		}
	}
}

func BGR2RGB(src gocv.Mat, dst *gocv.Mat) {
	split := gocv.Split(src)
	b, g, r := split[0], split[1], split[2]
	mergeSlice := []gocv.Mat{r, g, b}
	gocv.Merge(mergeSlice, dst)
}


func traverse(mat gocv.Mat) {
	var w = mat.Cols()
	var h = mat.Rows()
	var c = mat.Channels()
	fmt.Println(w, h, c)
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			fmt.Printf("%f,",mat.GetFloatAt3(x, y, 1))
		}
		fmt.Println()
	}
}
