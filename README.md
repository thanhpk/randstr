# Randstr

Randstr is an Go library for generate secure random strings

## Install
```
  go get -u github.com/thanhpk/randstr
```

## Usage
### Generate random hex string
```go
token := randstr.RandomHex(16) // generate 128-bit hex string
```
Running example
```go
  package main
  import(
    "github.com/thanhpk/randstr"
    "fmt"
  )
  
  func main() {
    for i := 0; i < 5; i++ {
      token := ranstr.RandomHex(16) // generate 128-bit hex string
      fmt.Println(token)
    }
  }
  // output
  //67aab2d956bd7cc621af22cfb169cba8
  //226eeb52947edbf3e97d1e6669e212c2
  //5f3615e95d103d14ffb5b655aa0eec1e
  //ff3ab4efbd74025b87b14b59422d304c
  //a6705813c174ca73ed795ea0bab12726
```

### Generate random ascii string
```go
token := randstr.RandomString(16) // generate a random 16 character length string
```
Running example
```go
  package main
  import(
    "github.com/thanhpk/randstr"
    "fmt"
  )
  
  func main() {
    for i := 0; i < 5; i++ {
      token := ranstr.RandomString(16)
      fmt.Println(token)
    }
  }
  // output
  // 7EbxkrHc1l3Ahmyr
  // I5XH2gc1EEHgbmGI
  // GlCycMpsxGkn9cDQ
  // U2OfBDQoak0z8FwV
  // kDX1m81u14YwEiCY
```
