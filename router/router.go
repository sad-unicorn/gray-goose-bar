package router

import "log"

func Dispatch(userId int, message string) {
	log.Println(userId, ":", message)
}