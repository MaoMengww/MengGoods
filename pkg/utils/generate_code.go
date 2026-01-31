package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

//生成六位验证码
func GenerateNumber (length int) (string, error) {
	if length <= 0 || length >18{
		return "", fmt.Errorf("输入的长度不合理")
	}
	min := int64(math.Pow10(length -1))
	max := int64(math.Pow10(length) - 1)
	rangeSize := max - min +1
	rangeBig := big.NewInt(int64(rangeSize))
	random, err := rand.Int(rand.Reader, rangeBig)
	if err != nil {
		return "", err
	}
	result := random.Int64() + min
	finalNumber := fmt.Sprintf("%d", result)
	return finalNumber, nil
}