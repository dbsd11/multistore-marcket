package main

import (
  "fmt"
  "github.com/manucorporat/try"
  "strconv"
  "encoding/json"
  "time"
)

type Goods struct {
  Id          string `json:"id"`
  name        string `json:"name"`
  code        string `json:"code"`
  Type        string `json:"type"`
  smallUrl    string `json:"smallUrl"`
  price       float64 `json:"price"`
  marcketName string `json:"marcketName"`
  inventory   int32 `json:"inventory"`
  tag         string `json:"tag"`
  star        string `json:"star"`
  createTime  time.Time `json:"createTime"`
}

func main()  {
  defer func() {
    fmt.Println("aaaa")
  }()

  var flag string = "G"
  fmt.Println(fmt.Sprintf("%s%06s", flag, strconv.Itoa(0)))


  try.This(func(){
    //panic("测试一些报错发生了")
    goodsBytes,_ := json.Marshal(&Goods{
      Id: "G00000",
      name: "棉衣",
      code:"my01",
      Type:"服装",
      smallUrl: "http://www.baidu.com",
      price: 100.0,
      marcketName: "华北01",
      inventory: 100,
      tag: "",
      star: "4.0",
      createTime: time.Now(),
    })
    fmt.Println(string(goodsBytes))
  }).Catch(func(e try.E) {
    // Print crash
    fmt.Println(e)
    recover()
  })

  recover()

  args := []string{"1", "2", "4"}
  fmt.Printf("Invoke function:%s args:%s \n", "aaaa", args)


  fmt.Println("adasdads")
}
