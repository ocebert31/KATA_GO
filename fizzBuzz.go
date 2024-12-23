package main

import "fmt"

func fizzBuzz() {
    var slice []interface{} 
    for i := 1; i < 100; i++ {
		if i % 5 == 0 && i % 3 == 0 {
			slice = append(slice, "fizzBuzz")
		} else if i % 5 == 0 {
			slice = append(slice, "buzz")
		} else if i % 3 == 0 {
            slice = append(slice, "fizz")
        } else {
            slice = append(slice, i) 
        }
    }
    fmt.Println(slice)
}