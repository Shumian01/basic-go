package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$6O9b9OJMM6gCnjiU9OWuXejXtuVIfr0n70DmyszKR0EjoNUYuhCEC"
	password := "xzl201515"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("❌ 不匹配:", err)
	} else {
		fmt.Println("✅ 匹配成功")
	}
}
