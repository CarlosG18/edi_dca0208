package main

import (
  
  "main/packs/list"
)

func main() {
  arraylist := list.ArrayList{}
  arraylist.Init(10)

  for i:=0; i<10; i++{
    arraylist.Add(i)
  }

  arraylist.RemoveOnIndex(200)
  arraylist.Remove()
  arraylist.AddOnIndex(59,2)
  arraylist.Set(1,100)

}
