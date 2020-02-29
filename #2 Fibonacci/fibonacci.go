package main

import "fmt"

func fibonacci(n int) int{
    if n < 3 {
      return 1
    }else{
        return fibonacci(n-1) + fibonacci(n-2)
    }
}

func main(){
  fmt.Println("Fibonacci RuleZ")

  for i:=1; i <= 10; i++ {
    fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
  }

}
