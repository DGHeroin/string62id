package String62Id

import (
    "crypto/md5"
    "fmt"
    "math/big"
    "strconv"
    "strings"
)
var chars = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")

//根据int64生成
func GetString62Int64(id int64) string {
    return getString62(encode62(id))
}

func GetString62BigInt(id *big.Int) string {
    return getString62(encode62Big(id))
}

func GetStringHash(str string) (result string) {
    hex := fmt.Sprintf("%x", md5.Sum([]byte(str)))
    res := make([]string, 4)
    for i := 0;i < 4;i++ {
        val, _ := strconv.ParseInt(hex[i*8:i*8+8], 16, 0)
        lHexLong := val & 0x3fffffff
        outChars := ""
        for j := 0;j < 6;j++ {
            outChars += chars[0x0000003D & lHexLong]
            lHexLong >>= 5
        }
        res[i] = outChars
    }
    result = strings.Join(res, "")
    return
}

func encode62(id int64) (result []int64) {
    for id > 0 {
        result = append(result, id % 62)
        id /= 62
    }
    return result
}

func encode62Big(id *big.Int) (result []int64) {
    zero := big.NewInt(0)
    big62 := big.NewInt(62)
    for id.Cmp(zero) == 1 {
        old := big.NewInt(0)
        val := old.Mod(id, big62).Int64()
        result = append(result, val)
        id.Div(id, big62)
    }
    return result
}

func getString62(indexA []int64) string {
    res := ""
    for _, val := range indexA {
        res += chars[val]
    }
    return reverseString(res)
}

// 反转字符串
func reverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes) - 1;from < to;from, to = from + 1, to - 1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}
