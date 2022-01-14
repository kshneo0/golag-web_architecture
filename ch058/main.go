package main

import (
	"context"
	"fmt"

	"github.com/web_archi/ch058/session"
)

// type key string
// var userKey key = "UserID"

type key int
var userKey key = 0
// var ipKey key = 1
// var isAdminKey key = 2
// var sessionKey key = 3

// func foo(ctx context.Context){
// 	ctx = context.WithValue(ctx, "currentFunc","foo")

//	// Log package may use this ctx value
// }

func main(){

	ctx := context.Background()
	ctx = session.WithUserID(ctx,1)

	userID := session.GetUserID(ctx)

	// ctx := context.WithValue(context.Background(),userKey, 1)
	//ctx := context.Background()

	// userID, ok := ctx.Value(userKey).(int)
	// userID, ok := ctx.Value("UserID").(int)
	// if !ok {
	if userID == nil {
		fmt.Println("Not Logged in!")
		return
	}
	fmt.Println(*userID)

	uId := ctx.Value(userKey)
	fmt.Println(uId)
	
}