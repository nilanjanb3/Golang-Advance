package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "nilanjan"

	hexCode := md5.Sum([]byte(text))
	res := hex.EncodeToString(hexCode[:])
	fmt.Println(res)
	fmt.Println(hex.DecodeString(res))

}
