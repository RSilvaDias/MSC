package main
import (
   "fmt"
)

// creating a function to multiply matrices
func MultiplyMatrix(matrixA [3][3]int, matrixB [3][2]int) [3][2]int {
   var total int = 0
   var result [3][2]int

   // multiplying matrices and storing result
   for i := 0; i < 3; i++ {
      for j := 0; j < 2; j++ {
         for k := 0; k < 3; k++ {
            total = total + matrixA[i][k]*matrixB[k][j]
         }
         result[i][j] = total
         total = 0
      }
   }
   return result
}
func main() {
   
   // initializing variables
   var result [3][2]int
   var i, j int
   matrixA := [3][3]int{
      {0, 1, 2},
      {4, 5, 6},
      {8, 9, 10},
   }
   matrixB := [3][2]int{
      {10, 11},
      {13, 14},
      {16, 17},
   }

   // printing matrices on the screen
   fmt.Println("The first matrix is:")
   for i = 0; i < 3; i++ {
      for j = 0; j < 3; j++ {
         fmt.Print(matrixA[i][j], "\t")
      }
      fmt.Println()
   }
   
   // printing a new line
   fmt.Println()
   fmt.Println("The second matrix is:")
   for i = 0; i < 3; i++ {
      for j = 0; j < 2; j++ {
         fmt.Print(matrixB[i][j], "\t")
      }
      fmt.Println()
   }
   fmt.Println()
   result = MultiplyMatrix(matrixA, matrixB)
   fmt.Println("The results of multiplication of matrix A & B: ")
   for i := 0; i < 3; i++ {
      for j := 0; j < 2; j++ {
         fmt.Print(result[i][j], "\t")
      }
      fmt.Println()
   }
}