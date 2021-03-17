/* 	break statement
The break statement in Go programming language has the following two usages −

	  When a break statement is encountered inside a loop, the loop is immediately
	terminated and the program control resumes at the next statement following the
	loop.

	  It can be used to terminate a case in a switch statement.

  If you are using nested loops, the break statement will stop the execution of
the innermost loop and start executing the next line of code after the block.
*/

package main

import "fmt"

func breakstatement() {
	breakstatementBasicOperation()
}

func breakstatementBasicOperation() int {
	var a int = 10

	for a < 20 {
		fmt.Printf("value of a: %d\n", a)
		a++
		if a > 15 {
			break
		}
	}
	return a
}
