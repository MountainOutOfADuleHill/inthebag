/*
June 20th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
*/

package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  reader := bufio.NewReader(os.Stdin)

  word, _ := reader.ReadString('\n')

  tilesInBag(word)

}

func tilesInBag(word string){
  letterMap := map[rune]int{'A': 9, 'B': 2, 'C': 2, 'D': 4, 'E': 12, 'F': 2,
    'G': 3, 'H': 2, 'I': 9, 'J': 1, 'K': 1, 'L': 4, 'M': 2, 'N': 6, 'O': 8, 'P': 2,
    'Q': 1, 'R': 6, 'S': 4, 'T': 6,'U': 4, 'V': 2, 'W': 2, 'X': 1, 'Y': 2, 'Z': 1, '_': 2}

  for _, letter := range word{
    for k, _ := range letterMap{//k is the key, v is the associated value
        if rune(letter) == k{
          letterMap[k] -= 1
        }
    }
  }

  for x, v := range letterMap{
    fmt.Println(string(x), ": ", v)
  }
}
