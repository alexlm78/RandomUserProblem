package main

import (
   "fmt"
   "time"
)

func main() {
   // Call the first solucion
   start := time.Now()
   mainOne()
   elapsed := time.Since(start)
   fmt.Println("Time elapsed for SolOne: ", elapsed)
   
   // Call the second solution
   start = time.Now()
   SolTwo()
   elapsed = time.Since(start)
   fmt.Println("Time elapsed for SolTwo: ", elapsed)
}
