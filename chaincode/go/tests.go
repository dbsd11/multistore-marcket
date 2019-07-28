package main

import (
  "fmt"
  "github.com/manucorporat/try"
  "strconv"
  "encoding/json"
  "time"
  "crypto/md5"
  "encoding/hex"
  "strings"
)

type Goods struct {
  Id          string `json:"id"`
  name        string `json:"name"`
  code        string `json:"code"`
  Type        string `json:"type"`
  smallUrl    string `json:"smallUrl"`
  Price       float64 `json:"price"`
  marcketName string `json:"marcketName"`
  inventory   int32 `json:"inventory"`
  tag         string `json:"tag"`
  star        string `json:"star"`
  createTime  time.Time `json:"createTime"`
}

type Trade struct {
  Id          string `json:id`
  GoodsName   string `json:"goodsName"`
  GoodsCode   string `json:"goodsCode"`
  MarcketName string `json:"marcketName"`
  CustomerNo  string `json:"customerNo"`
  TradePrice  float64 `json:"tradePrice"`
  TradeTime   time.Time `json:"tradeTime"`
  CreateTime  time.Time `json:"createTime"`
}

func main()  {
  defer func() {
    fmt.Println("aaaa")
  }()

  var tradeList = []Trade{}

  error := json.Unmarshal([]byte("[{\"customerId\":\"cus01\",\"goodsName\":\"棉衣\",\"marcketName\":\"华北01\"}]"), &tradeList)

  fmt.Println(error)

  tradeBytes,_ := json.Marshal(tradeList);

  fmt.Println(string(tradeBytes))

  fmt.Println()

  goods := Goods{
    Price:float64(100),
  }

  goods.Price += float64(100)*(float64(100000 - 83)/10000.0)

  goodsBytes,_ := json.Marshal(goods);
  fmt.Println(string(goodsBytes))

  var a int =0
  if a := 1; a!= 2{
    fmt.Println(a)
  }

  fmt.Println(a)



  fmt.Println(strings.Compare(Trade{Id:"1"}.Id, Trade{Id:"1"}.Id)==0)

  fmt.Println(getMd52("棉衣华北01"))

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
      Price: 100.0,
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


  fmt.Println(time.Now().Unix())


}

func getMd52(str string) string {
  h := md5.New()
  h.Write([]byte(str))
  return hex.EncodeToString(h.Sum(nil))
}
