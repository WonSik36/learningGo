package main

/* reference : http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80 */

// package

/* package import */
/* 
	Standard packages will be found in GOROOT / pkg and, in the case of user packages or 3rd party packages, in GOPATH / pkg.
*/

/* package scope */
/* 
	There are functions, structures, interfaces, and methods in the package. 
	If their names start with an uppercase letter, they can be used as public.
	On the other hand, if the name starts with a lowercase letter, 
	it is non-public and can only be used inside a package.
*/

/* package init */
/* 
	When you create a package, you can write an init () function that is called the first time the package is run. 
	That is, the init function is executed when the package is loaded. It is called automatically without calling.
*/

/* package init */
/* 
	When you create a package, you can write an init () function that is called the first time the package is run. 
	That is, the init function is executed when the package is loaded. It is called automatically without calling.
*/

/* package alias */
/* 
	In some cases, you might want to import a package and just call the init () function in that package. 
	In this case, an alias named _ is specified during package import.

	import _ "other/xlib"

	If the package names are the same, 
	but you want to load from different versions or different locations, 
	you can use package aliases to distinguish them.

	import (
    mongo "other/mongo/db"
    mysql "other/mysql/db"
	)
	func main() {
		mondb := mongo.Get()
		mydb := mysql.Get()
		//...
	}
*/

/* package build */
/*
	go install package file -> /pkg/windows_amd64
	ex) src/testlib> go install -> /pkg/windows_amd64/testlib.a

	go install main file -> /bin
*/

import "testlib"

func main() {
    song := testlib.GetMusic("Alicia Keys")
    println(song)
}