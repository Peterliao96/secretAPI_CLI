package main

import ("fmt"
"github.com/peterliao96/secret")

func main(){
  v := secret.Memory("fake-key")
  err := v.Set("demo_key", "well")
  if err != nil{
    panic(err)
  }

  plain, err := v.Get("demo_key")
  if err != nil{
    panic(err)
  }
  fmt.Println("Plain:",plain)
}
