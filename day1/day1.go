package day1

import ("fmt"
        "strings"
        "strconv"
        "io/ioutil")

func StartDay1(){
    
  data, err := ioutil.ReadFile("./day1/input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
  
  numberList := strings.Split(string(data),"\n")
    greaterCount := 0   
   for i:=3; i<len(numberList); i++ {   
     num1A, _ := strconv.ParseInt(numberList[i], 10, 64)
     num2A, _ := strconv.ParseInt(numberList[i-1], 10, 64)
     num3A, _ := strconv.ParseInt(numberList[i-2], 10, 64)
     
     num1B, _ := strconv.ParseInt(numberList[i-1], 10, 64)
     num2B, _ := strconv.ParseInt(numberList[i-2], 10, 64)
     num3B, _ := strconv.ParseInt(numberList[i-3], 10, 64)
     
     secondWindow := num1A + num2A + num3A
     firstWindow := num1B + num2B + num3B
     
     
     if (secondWindow > firstWindow){
       fmt.Println("sum ", secondWindow, " greater than ", firstWindow)
       greaterCount = greaterCount+1  
     }
   } 
  
     fmt.Println(greaterCount)
}
                      