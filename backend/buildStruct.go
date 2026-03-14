package main

import(
	"sync"
	"time"
)

type Build struct {
	ID int `json:"id"`
	Repo string `json:"repo"`
	Branch string `json:"branch"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}