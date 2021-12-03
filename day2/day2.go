package day2

import ("fmt"
        "regexp"
        "strings"
        "strconv"
        "io/ioutil")


func processInput() []string {
   data, err := ioutil.ReadFile("./day2/input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return make([]string, 0)
    }
  
  numberList := strings.Split(string(data),"\n")
  
  return numberList
}

func StartDay(){
  subDirections := processInput()
  
  aim := int64(0)
  moveForward := int64(0)
  moveUpDown := int64(0)
  
  for i :=0; i < len(subDirections); i++ {
    
    regexForward := "^forward"
    regexUp := "^up"
    regexDown := "^down"
    
    reForward := regexp.MustCompile(regexForward)
    findsForward := reForward.FindAllString(subDirections[i], 1)
    
    reUp := regexp.MustCompile(regexUp)
    findsUp := reUp.FindAllString(subDirections[i], 1)
    
    reDown := regexp.MustCompile(regexDown)
    findsDown := reDown.FindAllString(subDirections[i], 1)
    
    numberRegex := regexp.MustCompile("[0-9]+")
    findsDigits := numberRegex.FindAllString(subDirections[i], 1)
    
    if(len(findsDigits) == 1){
  
      num, _ := strconv.ParseInt(findsDigits[0], 10, 64)
        
      if(len(findsForward) == 1){
        moveForward += num
        newDepth := aim * num
        fmt.Println(subDirections[i], aim, num, newDepth)
        moveUpDown += newDepth
      }
      if(len(findsUp) == 1){
        //moveUpDown -= num
        aim -= num
      }
      if(len(findsDown) == 1){
        //moveUpDown += num
        aim += num
      }
    }
    
  }
  fmt.Println(moveUpDown)
  fmt.Println(moveForward)
  fmt.Println(moveUpDown * moveForward)
}