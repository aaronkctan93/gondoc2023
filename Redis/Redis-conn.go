package Redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// var redisdb *redis.Client

func main() {
	fmt.Println("golang連結redis")

	c := NewClinet()
	// test(c)
	SetTest(c, "A", "123")
	GetTest(c, "A")
	DelTest(c, "A")
	fmt.Println()
	HsetTest(c, "user_1", "username", "tizi365")

	HgetTest(c, "user_1", "username")
	HdelTest(c, "user_1", "username")
}

func NewClinet() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:30679",
		Password: "",
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

// func test(c *redis.Client) { // 對該 redis.Client 進行操作
// 	err := c.Set("key", "value", 0).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := c.Get("key").Result() // => GET key
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := c.Get("key2").Result() // => GET key2
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// }

func GetTest(c *redis.Client, key string) {
	val, err := c.Get(key).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println("Get Result:", val)

}

func SetTest(c *redis.Client, key, value string) {
	err := c.Set(key, value, 0).Err() // 0 意思是永遠不會過期
	if err != nil {
		panic(err)
	}
	fmt.Println("Set execution completed")
}

func DelTest(c *redis.Client, key string) {
	err := c.Del(key).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Del excetuion completed!")
}

func HsetTest(c *redis.Client, key, field, value string) {
	err := c.HSet(key, field, value).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGet execution completed!")
}

func HgetTest(c *redis.Client, key, field string) {
	field, err := c.HGet(key, field).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("HGet result:", field)
}

func HdelTest(c *redis.Client, key, field string) {
	c.HDel(key, field)

	fmt.Println("HDelete execution completed")
}
