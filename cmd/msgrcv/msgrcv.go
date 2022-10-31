// Message Receiver writes to redis
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	for {
		var tel string
		fmt.Print("Enter tel:")
		fmt.Scanf("%s", &tel)
		var msg string
		fmt.Print("Enter msg:")
		fmt.Scanf("%s", &msg)
		var phone string
		fmt.Print("Enter phone:")
		fmt.Scanf("%s", &phone)
		ctx := context.Background()
		err := rdb.Set(ctx, "M"+tel+"&"+msg, phone, time.Minute).Err()
		if err != nil {
			fmt.Println("Set failed")
		}
	}
}
