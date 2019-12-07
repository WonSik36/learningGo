package main
import "fmt"
 
// collection slice
// []int
// not [n]int

// slice has three value
// pointer to array
// length
// capacity

func main() {
	/* Nil slice */
	var a []int        // declare slice variable
	if a == nil {
        println("Nil Slice")
    }
	println(len(a),cap(a)) // length: 0, capacity: 0

    a = []int{1, 2, 3} // initialize literal value
    a[1] = 10
	fmt.Println(a)     // [1, 10, 3]
	
	/* make */
	s := make([]int, 5, 6)
	println(len(s), cap(s))	// length: 5, capacity: 6
	println(s[0], s[1], s[2], s[3], s[4])
	// println(s[5]) // runtime error
	
	s = append(s,6)
	println(s[5])
	println(len(s), cap(s))	// length: 6, capacity: 6

	s = append(s,7)
	println(s[6])
	println(len(s), cap(s))	// length: 7, capacity: 12

	/* sub slice */
	a = []int{0, 1, 2, 3, 4, 5}
	a = a[2:]  // [2 3 4 5]
	fmt.Println(a)
	a = a[0:1]  // [2]
	fmt.Println(a)
	
	/* append */

	var arr []int
	println(len(arr), cap(arr))
	arr = append(arr, 0,1,2,3)
	fmt.Println(arr)
	println(len(arr), cap(arr))

	/* append slice */

	sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
 
    sliceA = append(sliceA, sliceB...)
    //sliceA = append(sliceA, 4, 5, 6)
 
	fmt.Println(sliceA) // [1 2 3 4 5 6]
	
	/* copy slice */
	source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    fmt.Println(target)  // [0 1 2 ] 출력
    println(len(target), cap(target)) // 3, 6
}