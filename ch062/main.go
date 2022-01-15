package main

import (
	"context"
	"fmt"

	"github.com/web_archi/ch062/session"
)

func main() {
	ctx := context.Background()
	ctx = session.SetUserID(ctx,1)
	ctx = session.SetIsAdmin(ctx,true)
	i := session.GetUserID(ctx)
	b := session.GetIsAdmin(ctx)
	fmt.Printf("User %d is admin %t\n",i,b)
}