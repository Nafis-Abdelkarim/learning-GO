package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* Define an array containing programming languages */
	languages := [10]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "C#", "Go", "Rust",
	}

	/* Define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 5) // len(modern4)
	modern = languages[3:8]
	new := languages[8:10]

	fmt.Printf("classic labguages: %v\n", classics)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("new languages: %v\n", new)

	/*-----------------------------------*/
	allLangs := languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{
		"React", "Vu", "Angular", "Svelete",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4] // length 4 capacity 4
	frameworks = append(frameworks, "Meteor")
	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js framwroks: %v\n", jsFrameworks)
}
