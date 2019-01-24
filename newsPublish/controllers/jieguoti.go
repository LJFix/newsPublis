package main

import "fmt"

type student3 struct {
	name  string
	sex   string
	age   int
	score int
	addr  string
}

//结构体切片作为函数参数
func Demo2(arr []student3) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {

			//比较结构体成员
			//if arr[j].age < arr[j+1].age
			if arr[j].score > arr[j+1].score {

				//arr[j].name, arr[j+1].name = arr[j+1].name, arr[j].name

				//交换结构体数组中的元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	//fmt.Println(arr)
}
func main() {
	var stus []student3 = []student3{
		student3{"MC天赐", "男", 22, 88, "湖北荆州"},
		student3{"鱼尾纹", "男", 18, 798, "广东河源"},
		student3{"宝马", "男", 28, 8, "河南洛阳"}}

	//结构体切片作为函数参数是地址传递  引用传递  形参可以改变实参的值
	Demo2(stus)
	fmt.Println(stus)
}
