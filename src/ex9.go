package main

// collection array

func main(){
    var arr1 [3]int
	var arr2 = [...]int{1, 2, 3}
	var multiArr1 [3][4][5]int
	var multiArr2 = [2][3]int{
		{1,2,3},
		{4,5,6},
	}

    arr1[0] = 1
    arr1[1] = 2
    arr1[2] = 3
    println(arr1[1])

	println(arr2[2])
	// println(arr2[3])	// invalid statement

	println(multiArr1[0][1][2])	// 0
	multiArr1[0][1][2] = 10
	println(multiArr1[0][1][2])

	println(multiArr2[1][1])
}