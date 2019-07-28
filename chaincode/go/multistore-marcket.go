package main

import (
  "time"
  "encoding/json"
  "fmt"
  "strconv"
  "bytes"
  "crypto/md5"
  "encoding/hex"
  "strings"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  sc "github.com/hyperledger/fabric/protos/peer"
  "github.com/manucorporat/try"
)

type SmartContract struct {
}

type Goods struct {
  Id          string `json:id`
  Name        string `json:"name"`
  Code        string `json:"code"`
  Type        string `json:"type"`
  SmallUrl    string `json:"smallUrl"`
  Price       float64 `json:"price"`
  MarcketName string `json:"marcketName"`
  Inventory   int32 `json:"inventory"`
  Tag         string `json:"tag"`
  Star        string `json:"star"`
  CreateTime  time.Time `json:"createTime"`
}

type Marcket struct {
  Id          string `json:id`
  Name        string `json:"name"`
  Code        string `json:"code"`
  Address     string `json:"address"`
  CustomerNum int32 `json:"customerNum"`
  TradeAmount float64 `json:"tradeAmount"`
  Tag         string `json:"tag"`
  Star        string `json:"star"`
  CreateTime  time.Time `json:"createTime"`
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

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
  goodsList := []Goods{
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "棉衣",
      Code:"my01",
      Type:"服装",
      SmallUrl: "http://www.baidu.com",
      Price: 100.0,
      MarcketName: "华北01",
      Inventory: 10000,
      Tag: "男士",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "拖鞋",
      Code:"tx01",
      Type:"鞋子",
      SmallUrl: "http://www.baidu.com",
      Price: 10.0,
      MarcketName: "华北01",
      Inventory: 100000,
      Tag: "凉鞋、女士",
      Star: "4.1",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "真皮手表",
      Code:"watch01",
      Type:"手表",
      SmallUrl: "http://www.baidu.com",
      Price: 1000.0,
      MarcketName: "华北01",
      Inventory: 100,
      Tag: "真皮、手表",
      Star: "4.5",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "糕点",
      Code:"gd01",
      Type:"食品",
      SmallUrl: "http://www.baidu.com",
      Price: 1.0,
      MarcketName: "华北01",
      Inventory: 100000,
      Tag: "甜品",
      Star: "4.9",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "钳子",
      Code:"qz01",
      Type:"工具",
      SmallUrl: "http://www.baidu.com",
      Price: 25.0,
      MarcketName: "华南01",
      Inventory: 3000,
      Tag: "不锈钢",
      Star: "3.9",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "沙发",
      Code:"sf01",
      Type:"家具",
      SmallUrl: "http://www.baidu.com",
      Price: 5000.0,
      MarcketName: "华南01",
      Inventory: 2800,
      Tag: "舒适、懒人休息",
      Star: "3.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "老铁SUV",
      Code:"suv01",
      Type:"汽车",
      SmallUrl: "http://www.baidu.com",
      Price: 100000.0,
      MarcketName: "华南01",
      Inventory: 10,
      Tag: "省油",
      Star: "4.6",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "凌美swift",
      Code:"lm01",
      Type:"办公用品",
      SmallUrl: "http://www.baidu.com",
      Price: 2.0,
      MarcketName: "华东01",
      Inventory: 80000,
      Tag: "耐用",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "茶杯",
      Code:"cb01",
      Type:"百货",
      SmallUrl: "http://www.baidu.com",
      Price: 8.0,
      MarcketName: "华东01",
      Inventory: 100,
      Tag: "好看",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "G"),
      Name: "兄弟翡翠",
      Code:"xd01",
      Type:"珠宝",
      SmallUrl: "http://www.baidu.com",
      Price: 1000000.0,
      MarcketName: "华西01",
      Inventory: 1,
      Tag: "翡翠、稀有",
      Star: "5.0",
      CreateTime: time.Now(),
    },
  }

  marckets := []Marcket{
    {
      Id: getAndIncrId(APIstub, "M"),
      Name: "华北01",
      Code: "scn01",
      Address: "华北市场XX中心",
      CustomerNum: 0,
      TradeAmount: 0,
      Tag: "",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "M"),
      Name: "华南01",
      Code: "scs01",
      Address: "华南市场xx坊",
      CustomerNum: 0,
      TradeAmount: 0,
      Tag: "",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "M"),
      Name: "华东01",
      Code: "sce01",
      Address: "华东市场XX码头",
      CustomerNum: 0,
      TradeAmount: 0,
      Tag: "",
      Star: "4.0",
      CreateTime: time.Now(),
    },
    {
      Id: getAndIncrId(APIstub, "M"),
      Name: "华西01",
      Code: "scw01",
      Address: "华西XX市场",
      CustomerNum: 0,
      TradeAmount: 0,
      Tag: "",
      Star: "4.0",
      CreateTime: time.Now(),
    },
  }

  for _, goods := range goodsList {
    APIstub.PutState(goods.Id, Marshal(goods))
    APIstub.PutState(getMd5(goods.Name + goods.MarcketName), []byte(goods.Id))
  }

  for _, marcket := range marckets {
    APIstub.PutState(marcket.Id, Marshal(marcket))
    APIstub.PutState(getMd5(marcket.Name), []byte(marcket.Id))
  }
  return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
  // Retrieve the requested Smart Contract function and arguments
  function, args := APIstub.GetFunctionAndParameters()

  fmt.Printf("Invoke function:%s args:%s\n", function, args)

  // Route to the appropriate handler function to interact with the ledger appropriately
  if function == "initiateTrade" {
    return s.initiateTrade(APIstub, args)
  } else if function == "query" {
    return s.query(APIstub, args)
  } else if function == "demoClean" {
    return s.demoClean(APIstub, args)
  }

  return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initiateTrade(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  tradeObj := Trade{}

  json.Unmarshal([]byte(args[1]), &tradeObj)

  //校验商品是否合法
  var goods Goods
  if goodsIdBytes, error := APIstub.GetState(getMd5(tradeObj.GoodsName + tradeObj.MarcketName)); error != nil {
    return shim.Error("商品名不合法")
  } else {
    var goodsId = string(goodsIdBytes)
    goodsBytes, _ := APIstub.GetState(goodsId)

    goods = Goods{}
    json.Unmarshal(goodsBytes, &goods)
  }

  //校验市场是否合法
  var marcket Marcket
  if marcketIdBytes, error := APIstub.GetState(getMd5(tradeObj.MarcketName)); error != nil {
    return shim.Error("市场名称不合法")
  } else {
    var marcketId = string(marcketIdBytes)
    marcketBytes, _ := APIstub.GetState(marcketId)

    marcket = Marcket{}
    json.Unmarshal(marcketBytes, &marcket)
  }

  //校验交易合法性
  if (strings.Compare(goods.MarcketName, marcket.Name) != 0) {
    return shim.Error("市场中不存在该商品")
  }
  if (goods.Inventory <= 0) {
    return shim.Error("该商品库存不足,请尝试从其余市场购买")
  }

  //交易校验成功. 交易记录落地,扣减商品库存,增加市场销售额
  tradeObj.Id = getAndIncrId(APIstub, "T")
  tradeObj.TradePrice = goods.Price
  tradeObj.CreateTime = time.Now()
  APIstub.PutState(tradeObj.Id, Marshal(tradeObj))

  marcket.CustomerNum += 1
  marcket.TradeAmount += goods.Price
  APIstub.PutState(marcket.Id, Marshal(marcket))

  goods.Inventory -= 1;
  APIstub.PutState(goods.Id, Marshal(goods))

  //上浮下调对应商品类型的价格
  defer func() {
    fmt.Errorf("Error: failed to change goods price, tradeId:%s\n", tradeObj.Id)
    recover()
  }()

  try.This(func() {
    queryGoodsResponse := s.query(APIstub, []string{"goods"})
    var goodsList []Goods
    json.Unmarshal(queryGoodsResponse.GetPayload(), goods)

    var maxInvetoryGoods Goods
    maxInvetory := int32(0)
    var reduceTotalPrice float64 = 0.0
    for _, goodsThis := range goodsList {
      if strings.Compare(goodsThis.MarcketName, goods.MarcketName) != 0 || strings.Compare(goodsThis.Type, goods.Type) != 0 {
        continue
      }

      //1.上调交易商品的价格
      if strings.Compare(goodsThis.Id, goods.Id) == 0 {
        increasePrice := goodsThis.Price * (float64(100000 - goodsThis.Inventory) / 100000.0)
        goodsThis.Price += increasePrice
        goodsThis.Price = float64Price(goodsThis.Price)
        APIstub.PutState(goodsThis.Id, Marshal(goodsThis))
        reduceTotalPrice = increasePrice * float64(goodsThis.Inventory);
        continue
      }

      if goodsThis.Inventory > maxInvetory {
        maxInvetory = goodsThis.Inventory
        maxInvetoryGoods = goodsThis
      }
    }
    //2.下调库存量最大的商品价格
    maxInvetoryGoods.Price -= reduceTotalPrice / float64(maxInvetoryGoods.Inventory)
    maxInvetoryGoods.Price = float64Price(maxInvetoryGoods.Price)
    APIstub.PutState(maxInvetoryGoods.Id, Marshal(maxInvetoryGoods))
  }).Catch(func(e try.E) {
    fmt.Println(e)
  })

  return shim.Success([]byte(tradeObj.Id))
}

func (s *SmartContract) query(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }

  var flag string
  if (args[0] == "goods") {
    flag = "G"
  } else if (args[0] == "marckets") {
    flag = "M"
  } else if (args[0] == "trades") {
    flag = "T"
  } else {
    if (strings.HasPrefix(args[0], "debug")) {
      if debugBytes, error := APIstub.GetState(strings.TrimPrefix(args[0], "debug")); error != nil {
        fmt.Println(error)
      } else {
        fmt.Println(args[0] + " = " + string(debugBytes))
      }
    }
    return shim.Success([]byte("[]"))
  }

  startKey := fmt.Sprintf("%s%06s", flag, "0")
  endKey := fmt.Sprintf("%s%06s", flag, "999999")

  resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
  if err != nil {
    return shim.Error(err.Error())
  }
  defer resultsIterator.Close()

  // buffer is a JSON array containing QueryResults
  var buffer bytes.Buffer
  buffer.WriteString("[")

  bArrayMemberAlreadyWritten := false
  for resultsIterator.HasNext() {
    queryResponse, err := resultsIterator.Next()
    if err != nil {
      return shim.Error(err.Error())
    }
    // Add a comma before array members, suppress it for the first array member
    if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
    }

    buffer.WriteString(string(queryResponse.Value))
    bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")
  return shim.Success(buffer.Bytes())
}

func (s *SmartContract) demoClean(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  if (len(args) == 0) {
    return shim.Error("Denied")
  }
  if (strings.Compare(args[0], strconv.Itoa(int(time.Now().Unix() / 60))) != 0) {
    return shim.Error("Denied")
  }

  if error := APIstub.DelState(args[1]); error != nil {
    return shim.Error("Denied")
  }

  return shim.Success(nil)
}

func getAndIncrId(APIstub shim.ChaincodeStubInterface, flag string) string {
  defer func() {
    fmt.Println("getAndIncrId failed flag:" + flag)
    if error := recover(); error != nil {
      fmt.Println(error)
    }
  }()

  var globalIdKey = "global_id_" + flag
  var globalId int = 0
  try.This(func() {
    if globalIdBytes, error := APIstub.GetState(globalIdKey); error == nil {
      globalId, _ = strconv.Atoi(string(globalIdBytes))
      globalId += 1
    }
  }).Catch(func(e try.E) {
    // Print crash
    fmt.Println(e)
  })

  APIstub.PutState(globalIdKey, []byte(strconv.Itoa(globalId)))
  return fmt.Sprintf("%s%06s", flag, string(strconv.Itoa(globalId)))
}

func getMd5(str string) string {
  h := md5.New()
  h.Write([]byte(str))
  return hex.EncodeToString(h.Sum(nil))
}

func Marshal(v interface{}) []byte {
  if vBytes, error := json.Marshal(v); error != nil {
    return []byte{}
  } else {
    return vBytes
  }
}

func float64Price(price float64) float64 {
  if priceFloat64, error:= strconv.ParseFloat(fmt.Sprintf("%.2f", price), 64);error != nil{
    return float64(0)
  } else {
    return priceFloat64
  }
}

func main() {
  fmt.Println("multistore-marcket chaincode starting...")

  // Create a new Smart Contract
  err := shim.Start(new(SmartContract))
  if err != nil {
    fmt.Printf("Error creating new Smart Contract: %s", err)
  }

  fmt.Println("multistore-marcket chaincode started!")
}





