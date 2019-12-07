package main

// function


func main(){
	var message []byte = []byte{'H','e','l','l','o',' ','W','o','r','l','d'}
	var strMessage string = string(message)
	var multiString []string = []string{"hello","world","everyone"}

	sayRef(&strMessage)
	say(strMessage)
	sayMulti(multiString[0], multiString[1], multiString[2])

	println(sum(1,2,3,4))
	count, total := sum2(1,2,3,4)
	println(count,total)

	count, total = sum3(5,6,7,8)
	println(count,total)
}

func say(msg string){
	println(msg)
}

func sayRef(msg *string){
	println(*msg)
	*msg = "changed"
}

func sayMulti(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}

func sum(nums ...int) int{
	s := 0
	for _,n := range nums{
		s += n
	}
	return s
}

func sum2(nums ...int) (int, int){
	s := 0
	c := 0
	for _,n := range nums{
		s += n
		c++
	}
	return c,s
}

func sum3(nums ...int) (c int, s int){
	s = 0
	c = len(nums)
	for _,n := range nums{
		s += n
	}

	return
}