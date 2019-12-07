package main

// conditional statement

func main() {
    var k int = 1
	max := 100

	if k == 1 {
		println("One")
	} else if k == 2 {  //같은 라인
		println("Two")
	} else {   //같은 라인
		println("Other")
	}

	if val := k * 2; val > max {
		println(val)
	}else{
		println(val*2)
	}
}