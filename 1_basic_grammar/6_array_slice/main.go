package main

import (
	"fmt"
)

func main() {
	/////////////// [array] ///////////////
	var arr1 [5]int    //如果没有初始化，数组中所有变量都会被默认初始化为0(Go特性)
	arr1[0] = 1        //可以通过下标访问数组元素

	arr2 := [5]int{1, 2, 3, 4, 5}    		//数组也可以速记声明，还可以初始化
	fmt.Printf("[arr2]: %v\n", arr2)        //打印数组中所有元素

	arr_int1 := [...]int{1, 2, 3, 4, 5}
    arr_int2 := arr_int1
    arr_int2[1] = 22         //改变新数组中第二个元素的值
    fmt.Printf("[arr_int1]: %v\n", arr_int1)    //对新数组的操作不会对原数组有影响
    fmt.Printf("[arr_int2]: %v\n", arr_int2)    //新数组的值已经改变

    ChmodArray(arr_int1)
	fmt.Printf("[arr_int1]: %v\n", arr_int1)    //传参时也一样，对新数组的操作不会对原数组有影响
	fmt.Printf("[arr_int1 Length: %d]\n", len(arr_int1))
	
	mutil_arr := [3][2]string {
		{"2018MVP", "James Harden"},        //这里必须写逗号','
		{"2017MVP", "Russell Westbrook"},
		{"2016MVP", "Stephen Curry"},
	}

	for i, v := range mutil_arr {
		fmt.Printf("[mutil_arr %d]: %v\n", i, v)
	}



	/////////////// [slice] ///////////////
	fmt.Println("=====================[slice--声明]======================")
	var slice_int []int
	if slice_int == nil {
		fmt.Println("[slice_int]: [slice_int is empty]")
	}

	slice_byte := make([]byte, 5, 5)
	fmt.Printf("[slice_byte]: %v\n", slice_byte)

	slice_byte2 := make([]byte, 4, 5)
	fmt.Printf("[slice_byte2]: %v\n", slice_byte2)

	slice_int2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("[slice_int2]: %v\n", slice_int2)

	arr_string := [5]string{"Alston", "T_MAC", "Battier", "Scola", "Yao"}
	slice_string := arr_string[0:3]
	slice_string2 := arr_string[0:5]
	fmt.Printf("[slice_string]: %v\n", slice_string)
	fmt.Printf("[slice_string2]: %v\n", slice_string2)

	fmt.Println("=====================[slice--声明]======================")

	fmt.Println("=====================[slice--操作]======================")
	fmt.Printf("slice_string[2]: %s\n", slice_string[2])
	// fmt.Printf("slice_string[3]: %s\n", slice_string[3])		//错误，下标越界

	var slice_string3 []string
	slice_string4 := []string{"Python", "Java"}
	slice_string3 = append(slice_string, "C++", "Go", "Rust")	//添加任意数量的元素
	slice_string3 = append(slice_string, slice_string4...)		//添加切片，注意切片后面一定要写...
	fmt.Printf("[slice_string3]: %v\n", slice_string3)

	num_slice := []int{1, 2, 3, 4, 5}
	ChmodSlice(num_slice)
	fmt.Printf("[num_slice]: %v\n", num_slice)		//输出结果：[num_slice]: [1 22 3 4 5]
	
	fmt.Println("=====================[slice--操作]======================")


	fmt.Println("=====================[slice--原理]======================")
	//当切片长度已达到容量，再添加元素
	car_array := [5]string{"BMW", "Benz", "Audi", "Toyota", "Honda"}
	car_slices := car_array[:]
	car_slices[4] = "KIA"
	fmt.Printf("[car_array]: %v\n", car_array)		//切片car_slices引用原数组car_array，对切片的操作也会影响到原数组
	fmt.Printf("[car_slices (cap : %d)]: %v\n", cap(car_slices), car_slices)	//输出结果：[car_slices (cap : 5)]: [BMW Benz Audi Toyota Honda]

	car_slices = append(car_slices, "Cadillac")
	fmt.Printf("[car_slices (cap : %d)]: %v\n", cap(car_slices), car_slices)	//输出结果：[car_slices (cap : 10)]: [BMW Benz Audi Toyota Honda Cadillac]

	car_slices[4] = "Volvo"
	fmt.Printf("[car_array]: %v\n", car_array)		//切片引用了新的数组，对切片的操作不会影响到原数组car_array

	//当切片长度未到容量时，添加元素
	phone_array := [...]string{"iPhone", "samsung", "NOKIA", "HUAWEI", "XIAOMI", "Vivo", "Oppo"}
	var phone_slices []string = phone_array[1:5]
	fmt.Printf("[phone_array]: %v\n", phone_array)  	//输出结果为：[iPhone samsung NOKIA HUAWEI XIAOMI Vivo Oppo]

	phone_slices = append(phone_slices, "MEIZU")
	fmt.Printf("[phone_array]: %v\n", phone_array)  	//输出结果为：[iPhone samsung NOKIA HUAWEI XIAOMI MEIZU Oppo]
	fmt.Println("=====================[slice--原理]======================")
}

func ChmodArray(arr_int [5]int) {
    arr_int[1] = 22
}

func ChmodSlice(num_slice []int) {
	num_slice[1] = 22
}