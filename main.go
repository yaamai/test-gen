package main
import (
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

func main() {
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
