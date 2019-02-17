package main

import "secret"

func main(){
  v := secret.Memory("fake-key")
  err := v.set("demo_key", "well")
  if err != nil{
    panic(err)
  }

  plain, err := v.Get("demo_key")
  if err != nil{
    panic(err)
  }
  fmt.Println("Plain:",plain)
}
