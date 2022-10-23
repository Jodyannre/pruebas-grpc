package main1



import (
	"fmt"
	//"github.com/go-redis/redis/v9"
	"context"
)



/*	Para Grcp
1) Instalar el proto 
	sudo apt install protobuf-compiler

*/

func main1(){

}





//Instalar dependencia para utilizar redis para version 7 de redis go get github.com/go-redis/redis/v9
func redis(){
	var ctx = context.Background()
	ip := "172.17.0.2:6379"
	rdb := redis.NewClient(&redis.Options{
		Addr: ip,
		Password:"",
		DB:0,
	})
	err := rdb.Set(ctx, "key", "value",0).Err()
	if err != nil{
		panic(err)
	}
	val,err := rdb.Get(ctx, "key").Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("key",val)
}
