# Randstr

Randstr is an Go library for generate random secure string

## Install
```
  go get github.com/thanhpk/randstr
```

## Usage example

```go
  package main
  import(
    "github.com/thanhpk/randstr"
    "fmt"
  )
  
  func main() {
    for i:=0; i< 10; i++ {
      token, _ := ranstr.RandomHex(16) // generate 128-bit hex string
      fmt.Println(token)
    }
  }
  // output
  //67aab2d956bd7cc621af22cfb169cba8
  //226eeb52947edbf3e97d1e6669e212c2
  //5f3615e95d103d14ffb5b655aa0eec1e
  //ff3ab4efbd74025b87b14b59422d304c
  //2d663439f4c5a5fc9bb816c9f20ccb89
  //393d4bff9705b8f50c100fdc808d5e60
  //3376b9a8d9696bc05431b3f32c3e314d
  //3e8ce977678a5840b63780ecba3e5b77
  //023d48b76c8a7f8c7f30d792316f1a64
  //a6705813c174ca73ed795ea0bab12726
  
```
