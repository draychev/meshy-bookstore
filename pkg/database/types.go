package database

import "github.com/draychev/go-toolbox/pkg/logger"

const (
	KeyTotalBooks = "total-books"
)

// Record stores key value pairs
type Record struct {
	Key      string `gorm:"primaryKey"`
	ValueInt int64
}

var log = logger.NewPretty("datbase")
