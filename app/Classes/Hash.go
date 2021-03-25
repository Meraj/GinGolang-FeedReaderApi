package Classes

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Hash struct {

}
func (Hash) HashString(pwd string) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func (Hash) CompareHashes(text string, hash string) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(text))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (Hash) GenerateToken() string {
	t := time.Now().Unix()
	token := make([]byte, 50)
	rand.Read(token)
	return fmt.Sprintf("%x"+string(t), token)
}