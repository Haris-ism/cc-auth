package utils

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func GenerateRandom(length int)(int,error){
	num:="1234567890"
	buffer := make([]byte, length)
    _, err := rand.Read(buffer)
    if err != nil {
        fmt.Println("ieu err")
		return 0,err
    }
	for i := 0; i < length; i++ {
        buffer[i] = num[int(buffer[i])%length]
    }
	result,err:=strconv.Atoi(string(buffer))
	return result,nil
}