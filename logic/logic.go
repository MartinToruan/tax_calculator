package logic

import "time"

type TaxLogic struct {
	MAX_CONCURRENT int
	connPool       chan struct{}
}

func (l *TaxLogic) Init() {
	l.connPool = make(chan struct{}, l.MAX_CONCURRENT)
	for i := 0; i < l.MAX_CONCURRENT; i++ {
		l.connPool <- struct{}{}
	}
}

func (l *TaxLogic) DeInit() {
	// Let The Logic complete the remaining Task
	time.Sleep(15 * time.Second)

	// Close Connection Pool
	close(l.connPool)

	// Draining the Channel
	for _ = range l.connPool {
	}
}
