package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "regexp"
  "math/big"
)

type Item struct {
  worrylv *big.Int
}

type Operation struct {
  first string
  sign string
  second string
}

type Queue struct {
  items []Item
}

func (q *Queue) Insert(item Item) {
  q.items = append(q.items, item) 
}

func (q *Queue) Len() int {
  return len(q.items)
}

func (q * Queue) Pop() Item {
  removeditem := q.items[len(q.items) - 1]
  q.items = q.items[0:len(q.items) - 1]
  return removeditem
}

type Monkey struct {
  items Queue 
  operation Operation 
  test *big.Int
  iftrue *Monkey
  iffalse *Monkey
  inspectioncount int
}

type ThrowRel struct {
  iftrue int
  iffalse int
}

func (m *Monkey) InspectElements() {
  for m.items.Len() > 0 {
    m.inspectioncount++
    item := m.items.Pop()
    newworrylv := m.CalcNewWorry(item)
    // divides by three and rounds down to the nearest integer
    //item.worrylv = new(big.Int).Div(newworrylv, big.NewInt(int64(3)))
    item.worrylv = new(big.Int).Mod(newworrylv, big.NewInt(int64(96577)))
    //item.worrylv = newworrylv
    worryTest := new(big.Int).Mod(item.worrylv, m.test)
    var monkeytothrow *Monkey
    if worryTest.BitLen() == 0 {
      monkeytothrow = m.iftrue
    } else {
      monkeytothrow = m.iffalse
    }
    monkeytothrow.items.Insert(item)
  }
}

func (m *Monkey) CalcNewWorry(item Item) *big.Int {
  oldworry := item.worrylv 
  result := big.NewInt(0)
  if m.operation.sign == "+" {
    if m.operation.second == "old" {
      result = oldworry.Add(oldworry, oldworry)
    } else {
      result = oldworry.Add(oldworry, ConvertStrToBigInt(m.operation.second))
    }
  } else if m.operation.sign == "*" {
    if m.operation.second == "old" {
      result = oldworry.Mul(oldworry, oldworry)
    } else {
      result = oldworry.Mul(oldworry, ConvertStrToBigInt(m.operation.second))
    }
  } else {
    fmt.Println("error m.operation.sign not reconized!!!")
  }
  return result
}

func ParseOperation(operation string) Operation {
  strsplit := strings.Split(operation, " ")
  return Operation{
    first: strsplit[0],
    sign: strsplit[1],
    second: strsplit[2],
  }
}

func ConvertStrToBigInt(str string) *big.Int {
  number, err := strconv.Atoi(string(str))
  if err != nil {
    fmt.Println("Error during conversion")
    return big.NewInt(-1)
  }
  return big.NewInt(int64(number))
}

func CalcOperation(operation string, oldworry *big.Int) *big.Int {
  strsplit := strings.Split(operation, " ")
  opsign := strsplit[1]
  var op2 string = strsplit[2]
  var result *big.Int = big.NewInt(0)
  if opsign == "+" {
    if op2 == "old" {
      result = oldworry.Add(oldworry, oldworry)
    } else {
      charN, err := strconv.Atoi(string(op2))
      if err != nil {
        fmt.Println("Error during conversion")
        return big.NewInt(-1)
      }
      result = oldworry.Add(oldworry, big.NewInt(int64(charN)))
    }
  } else if opsign == "*" {
    if op2 == "old" {
      result = oldworry.Mul(oldworry, oldworry)
    } else {
      charN, err := strconv.Atoi(string(op2))
      if err != nil {
        fmt.Println("Error during conversion")
        return big.NewInt(-1)
      }
      result = oldworry.Mul(oldworry, big.NewInt(int64(charN)))
    }
  } else {
    fmt.Println("error opsign not reconized!!!")
  }
  return result
}

func main() {
  readFile, err := os.Open("input")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  monkeys := []*Monkey{}
  var startingitemsreg = regexp.MustCompile("Starting items:")
  var operationreg= regexp.MustCompile("Operation:")
  var testreg = regexp.MustCompile("Test:")
  var iftruereg = regexp.MustCompile("If true:")
  var iffalsereg = regexp.MustCompile("If false:")
  var monkeyreg = regexp.MustCompile(`Monkey \d`)
  var m * Monkey
  var trel [] ThrowRel 
  for fileScanner.Scan() {
    str := fileScanner.Text()
    // line Starting items:
    if monkeyreg.Match([]byte(str)) {
      var monkey Monkey 
      m = &monkey
    }
    if startingitemsreg.Match([]byte(str)) {
      strsplit := strings.Split(str, "Starting items:")
      cnumbers := strings.Split(strsplit[1], ",")
      for i := 0; i < len(cnumbers); i++ {
        tnum := strings.Trim(cnumbers[i], " ")
        num, err := strconv.Atoi(string(tnum))
        if err != nil {
          fmt.Println("Error during conversion")
          return
        }
        m.items.Insert(Item{
          worrylv: big.NewInt(int64(num)),
        })
      }
    }
    // match operation
    if operationreg.Match([]byte(str)) {
      strsplit := strings.Split(str, "new = ")
      m.operation = ParseOperation(strsplit[1])
    }
    // test
    if testreg.Match([]byte(str)) {
      strsplit := strings.Split(str, "divisible by ")
      intN, err := strconv.Atoi(string(strsplit[1]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      m.test = big.NewInt(int64(intN))
    }
    // if true
    if iftruereg.Match([]byte(str)) {
      strsplit := strings.Split(str, "throw to monkey ")
      intN, err := strconv.Atoi(string(strsplit[1]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      trel = append(trel, ThrowRel{
        iftrue: intN,
      })
    }

    // if false
    if iffalsereg.Match([]byte(str)) {
      strsplit := strings.Split(str, "throw to monkey ")
      intN, err := strconv.Atoi(string(strsplit[1]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      trel[len(trel) - 1].iffalse = intN
      monkeys = append(monkeys, m)
    }
  }

  for i, _ := range trel {
    monkeys[i].iftrue = monkeys[trel[i].iftrue]   
    monkeys[i].iffalse = monkeys[trel[i].iffalse]   
  }

  for _, m := range monkeys {
    fmt.Println(m)
  }

  // go through 10k rounds
  for i := 0; i < 10000; i++ {
    // go through all the monkeys
    for _, monkey := range monkeys {
      monkey.InspectElements()  
    }
  }
  for _, m := range monkeys {
    fmt.Println(m.inspectioncount)
  }
}
