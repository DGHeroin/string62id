package String62Id

import (
    "fmt"
    "log"
    "math/big"
    "testing"
    "time"
)

var (
    testValue = int64(1234567890)
)

func TestGetString62Simple(t *testing.T) {
    for i := int64(0); i < 10; i++ {
        fmt.Printf("int64: %v => %s\n", i, GetString62Int64(i))
    }
}

func TestGetString62Int(t *testing.T) {
    fmt.Printf("int64: %s\n", GetString62Int64(testValue))
}


func TestGetString62BigInt(t *testing.T) {
    fmt.Printf("big.Int: %s\n", GetString62BigInt(big.NewInt(testValue)))
}

func TestGetString62BigIntSuperLarge(t *testing.T) {
    val := big.NewInt(0)
    // 3 times md5 hash string of "hello-world"
    val.SetString("HjmOjqKTnOveG91nuuS9L9jeHjmOjqKTnOveG91nuuS9L9jeHjmOjqKTnOveG91nuuS9L9je", 62) // set a super large number

    fmt.Printf("large big.Int: %v => %s\n", val.Text(10), GetString62BigInt(val))
}

func TestStringHash(t *testing.T) {
    fmt.Printf("string hash: %s\n", GetStringHash("hello-world"))
}

func TestBigIntAdd(t *testing.T)  {
    num := big.NewInt(0)
    isRunning := true
    var start time.Time
    go func() {
        start = time.Now()
        for isRunning {
            num.Add(num, big.NewInt(1))
        }
        val := num.Int64()
        elapsed := time.Now().Sub(start)
        ops := elapsed.Nanoseconds() / val
        opPerSec := float64(val) / elapsed.Seconds()
        log.Printf("done: %v time:%v %v ns/ops %f op/s\n", val,  elapsed, ops, opPerSec)
    }()
    select {
    case <- time.After(time.Second * 10):
        isRunning = false
    }
}