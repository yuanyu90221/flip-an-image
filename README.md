# flipping-an-image

### 題目來源:
[flipping-an-image](https://leetcode.com/problems/flipping-an-image/)

### 原文:

Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
### 解讀：

給定 二元 矩陣 A,

目標是要把矩陣 做兩個操作

第一個是 flip 就是讓原本每個 A[i][j] = A[i][len(A[i]) -1 -j]

第二個是 invert 就是讓原本每個元素值 原本是0變成1
                                原本是1變成0

## 初步解法:
### 初步觀察:
首先是這兩個操作可以分開來看

flip 操作是針對位置關係做調換 也就是 假設給定某個i row

對於每個 j >= 0 , j < len(A[i]), B[i][j] = A[i][len(A) -1 -j]

invert 操作 是針對每個值做操作 也就是新的 值 是舊的值 +1 % 2

B[i][j] = (B[i][j] + 1)%2

把這兩個操作 合成一個操作就編成

C[i][j] = (A[i][len(A[i])-1-j] + 1)%2

參考別人的作法

其實發現  flip左右兩邊 對稱的位置值互換

所以可以只做一半的loop

### 初步設計:
Given a binary matrix A,

Step 0: let idxY = 0, idxX = 0, sameSize of matrix B

Step 1: if idxY > len(A) go to step 6

Step 2: if idxX > len(A[idxY]) go to step 5

Step 3:   B[idxY][idxX] = (A[idxY][len(A[idxY]) - 1 - idxX] + 1) %2

Step 4: idxX += 1 go to step 2

Step 5: idxY += 1 go to step 1

Step 6: return B

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
最初想法
```golang
func flipAndInvertImage(A [][]int) [][]int {
    lenY := len(A)
	lenX := len(A[0])
	result := make([][]int, lenY)
	for idxY := 0; idxY < lenY; idxY++ {
        result[idxY] = make([]int, lenX)
		for idxX := 0; idxX < lenX; idxX++ {
			result[idxY][idxX] = (A[idxY][lenX-1-idxX] + 1) % 2
		}
	}

	return result
}
```
參考別人的作法
```golang
package flip_image

func flipAndInvertImage(A [][]int) [][]int {
	lenY := len(A)
	lenX := len(A[0])
	halfX := lenX / 2
	if lenX%2 == 1 {
		halfX += 1
	}
	for idxY := 0; idxY < lenY; idxY++ {
		for idxX := 0; idxX < halfX; idxX++ {
			A[idxY][idxX], A[idxY][lenX-1-idxX] = (A[idxY][lenX-1-idxX]+1)%2, (A[idxY][idxX]+1)%2
		}
	}
	return A
}

```
## 測資的撰寫

```golang
package flip_image

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				A: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name: "Example2",
			args: args{
				A: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			},
			want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 30th challenge](https://hackmd.io/hlfS-iMrTAmH07I7veOfow?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)