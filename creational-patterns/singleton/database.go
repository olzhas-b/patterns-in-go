package singleton

import (
	"log"
	"sync"
)

type DB struct {
	conn *MyConn
	once sync.Once
}

type MyConn struct {
	DNS string
}

var db DB

func GetDatabaseConn(dns string) *MyConn {
	db.once.Do(func() {
		//dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env["USER"], env["PASSWORD"], env["HOST"], env["PORT"], env["NAME"])
		//ctx := context.Background()
		//conn, err := pgxpool.Connect(ctx, dns)
		db.conn = &MyConn{DNS: dns}
		log.Printf("connected to %s", dns)
	})

	return db.conn
}
