package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "syscall"
  "golang.org/x/crypto/bcrypt"
  "golang.org/x/crypto/ssh/terminal"
)
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
func accountInput() (string, string){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter Username: ")
  username, _ := reader.ReadString('\n')
  fmt.Print("Enter Password: ")
  bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
  if err != nil {
    fmt.Println("An error occurred")
  }
  password := string(bytePassword)
  return strings.TrimSpace(username), strings.TrimSpace(password)
}
func main() {
  _, pw := accountInput()
  hashpw,_ := hashPassword(pw)
  check := checkPasswordHash(pw, hashpw)
  if !check{
    fmt.Println("An error occured")
  }
  fmt.Println("hash")
  fmt.Println(hashpw)
}
