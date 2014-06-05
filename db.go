package unicore_network

import (
	"github.com/fzzy/radix/redis"
	"github.com/go-martini/martini"
	"os"
	"strconv"
	"time"
)

/*
NOTE: http://blog.gopheracademy.com/day-11-martini explains how this stuff
      mostly works. It's a bit like Django middleware but allows one to
      insert extra parameters that are passed along to the handler (view)
*/
func DB() martini.Handler {
	// connect to the db
	redis_network := os.Getenv("REDIS_NETWORK")
	redis_addr := os.Getenv("REDIS_ADDR")
	redis_db := os.Getenv("REDIS_DB")
	if redis_network == "" {
		redis_network = "tcp"
	}
	if redis_addr == "" {
		redis_addr = "127.0.0.1:6379"
	}
	if redis_db == "" {
		redis_db = "0"
	}

	// NOTE: On a low level this uses net.Dial, see:
	//		 http://golang.org/pkg/net/#Dial for details on network & addr
	conn, err := redis.DialTimeout(
		redis_network, redis_addr,
		time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	// close the connection when done
	defer conn.Close()

	// select db
	redis_db_number, err := strconv.Atoi(redis_db)
	if err != nil {
		panic(err)
	}
	conn.Cmd("select", redis_db_number)

	return func(c martini.Context) {
		// make available to subsequent handlers
		c.Map(conn)
		c.Next()
	}
}
