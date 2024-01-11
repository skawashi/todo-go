package model

type Todo struct {
  ID          uint `gorm:"primary_key"`
  Task        string
  Description string
}
