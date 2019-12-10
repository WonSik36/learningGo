package main
 
import (
	"log"
	"os"
)

// error

type MyError struct {
	message string
}

func (e MyError) Error() string{
	return e.message
}

 
func main() {
    f, err := os.Open("./src/ex15.go")
    if err != nil {
        log.Fatal(err.Error())
    }
	println(f.Name())
	
	err = MyError{message : "my error"}

	switch err.(type) {
		default: // no error
			println("ok")
		case MyError:
			log.Print(err.Error())
		case error:
			log.Fatal(err.Error())
		}
}