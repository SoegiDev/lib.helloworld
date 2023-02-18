package greetings

import ("testing"
	"fmt"
)

// function to test if the original
// go program is working
func TestGreetings(test *testing.T){
  
      expected:= "Hi, Fajar. Welcome!"
	  ret := Hello("Fajar")

	  if ret == expected{
		fmt.Printf("Successful!\n")
	  }else{
		test.Errorf("Error!\n")
	  }
	 
  
}