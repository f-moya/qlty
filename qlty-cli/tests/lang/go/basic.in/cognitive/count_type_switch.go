package foo

import "fmt"

func foo() {
		var x interface{}

		switch i := x.(type) {                 // +1
			case nil:
					fmt.Printf("type of x :%T",i)
			case int:
					fmt.Printf("x is int")
			case float64:
					fmt.Printf("x is float64")
			case func(int) float64:
					fmt.Printf("x is func(int)")
			case bool, string:
					fmt.Printf("x is bool or string")
			default:
					fmt.Printf("don't know the type")
		}
}
