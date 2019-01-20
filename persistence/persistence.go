package persistence

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Persistence is a module that communicate with Database
// Every request to Database must through this Module

// Contains Database Infomation and Connection Pool
type TaxPersistence struct {
	DB_NAME       string
	DB_HOST       string
	DB_PORT       int
	DB_USER       string
	DB_PASSWORD   string
	MAX_DB_CLIENT int
	dbClientPool  chan *sql.DB
}

// This function is used to setup the Persistence.
// This function will generate and set the connection pool
func (p *TaxPersistence) Init() {
	// Init Connection Pool
	p.dbClientPool = make(chan *sql.DB, p.MAX_DB_CLIENT)

	// Construct Database Info
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=3 sslmode=disable", p.DB_HOST, p.DB_PORT, p.DB_USER, p.DB_PASSWORD, p.DB_NAME)

	for i := 0; i < p.MAX_DB_CLIENT; i++ {
		// Create Pool Connection
		db, err := sql.Open("postgres", dbInfo)
		if err != nil {
			fmt.Println("can't connect to database: ", dbInfo)
		}

		// Insert Client to Pool
		p.dbClientPool <- db
	}
}

// This function is used to Turn Off the Persistence.
func (p *TaxPersistence) DeInit() {
	// Let Persistence to complete the remaining work
	time.Sleep(15 * time.Second)

	// Close Channel Pool
	close(p.dbClientPool)

	// Draining the Channel and Close all DBClient
	for db := range p.dbClientPool {
		db.Close()
	}
}

// This function is used to set dbClientPool
// Current usage is for unit test
func (p *TaxPersistence) SetClientPool(clients chan *sql.DB) {
	p.dbClientPool = clients
}
