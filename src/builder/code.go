package main

import (
	"fmt"
	"time"
)

type DBPool struct {
	dsn             string
	maxOpenConn     int
	maxIdleConn     int
	maxConnLifeTime time.Duration
}

type DBPoolBuilder struct {
	DBPool
	err error
}

func Builder () *DBPoolBuilder {
	b := new(DBPoolBuilder)
	// 设置 DBPool 属性的默认值
	b.DBPool.dsn = "127.0.0.1:3306"
	b.DBPool.maxConnLifeTime = 1 * time.Second
	b.DBPool.maxOpenConn = 30
	return b
}

func (b *DBPoolBuilder) DSN(dsn string) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if dsn == "" {
		b.err = fmt.Errorf("invalid dsn, current is %s", dsn)
	}

	b.DBPool.dsn = dsn
	return b
}

func (b *DBPoolBuilder) MaxOpenConn(connNum int) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if connNum < 1 {
		b.err = fmt.Errorf("invalid MaxOpenConn, current is %d", connNum)
	}

	b.DBPool.maxOpenConn = connNum
	return b
}

func (b *DBPoolBuilder) MaxConnLifeTime(lifeTime time.Duration) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if lifeTime < 1  * time.Second {
		b.err = fmt.Errorf("connection max life time can not litte than 1 second, current is %v", lifeTime)
	}

	b.DBPool.maxConnLifeTime = lifeTime
	return b
}

func (b *DBPoolBuilder) Build() (*DBPool, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.DBPool.maxOpenConn < b.DBPool.maxIdleConn {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.DBPool.maxOpenConn, b.DBPool.maxIdleConn)
	}
	return &b.DBPool, nil
}

func main() {
	dbPool, err := Builder().DSN("localhost:3306").MaxOpenConn(50).MaxConnLifeTime(10 * time.Second).Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbPool)
}
