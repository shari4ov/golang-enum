package main

import (
	"fmt"
	"golang-enum/status"
)

func main() {
	updateUserStatus(1, status.Success)
}

func updateUserStatus(userId int, status status.Status) {
	fmt.Printf("user_id %d,status %s", userId, status)
}
