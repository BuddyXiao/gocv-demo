/*
 * @Author: 肖和锋
 * @Date: 2022/7/7 14:43
 */
package examples

import (
	"fmt"
	"testing"
)

func TestClassify(t *testing.T) {
	var tests = []struct {
		imgPath string
		real string
	}{
		{"config/baihe.jpg", "baihe"},
		{"config/dangshen.jpg","dangshen"},
		{"config/gouqi.jpg","gouqi"},
		{"config/huaihua_12.jpg","huaihua"},
		{"config/jinyinhua.jpg","jinyinhua"},
	}
	for i := 0; i < len(tests); i++ {
		pred := classify(tests[i].imgPath)
		real := tests[i].real
		if pred != real {
			t.Errorf("[%v] predict:%s  real:%s,  tests fail pass \n", i+1, pred,real)
		}
	}
}

func TestOne(t *testing.T) {
	s := classify("config/baihe.jpg")
	fmt.Println(s)
}
