package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
)

// 数据导出器
type IDataExporter interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string, writer io.Writer) error
}

// 数据查询器
type IDataFetcher interface {
	Fetch(sql string) []interface{}
}

type MysqlDataFetcher struct {
	Config string
}

func (mf *MysqlDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source: " + mf.Config)
	rows := make([]interface{}, 0)
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

func NewMysqlDataFetcher(configStr string) IDataFetcher {
	return &MysqlDataFetcher{
		Config: configStr,
	}
}

type OracleDataFetcher struct {
	Config string
}

func NewOracleDataFetcher(configStr string) IDataFetcher {
	return &OracleDataFetcher{
		configStr,
	}
}

func (of *OracleDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from oracle source: " + of.Config)
	rows := make([]interface{}, 0)
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}



type CsvExporter struct {
	mFetcher IDataFetcher
}

func NewCsvExporter(fetcher IDataFetcher) IDataExporter {
	return &CsvExporter{
		fetcher,
	}
}

func (ce *CsvExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *CsvExporter) Export(sql string, writer io.Writer) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
	for i, v:= range rows {
		fmt.Printf("  行号: %d 值: %s\n", i + 1, v)
	}
	return nil
}

type JsonExporter struct {
	mFetcher IDataFetcher
}

func NewJsonExporter(fetcher IDataFetcher) IDataExporter {
	return &JsonExporter{
		fetcher,
	}
}

func (je *JsonExporter) Fetcher(fetcher IDataFetcher) {
	je.mFetcher = fetcher
}

func (je *JsonExporter) Export(sql string, writer io.Writer) error {
	rows := je.mFetcher.Fetch(sql)
	fmt.Printf("JsonExporter.Export, got %v rows\n", len(rows))
	for i, v:= range rows {
		fmt.Printf("  行号: %d 值: %s\n", i + 1, v)
	}
	return nil
}


func main() {
	mFetcher := NewMysqlDataFetcher("mysql://127.0.0.1:3306")
	csvExporter := NewCsvExporter(mFetcher)
	var writer bytes.Buffer
	// 从MySQL数据源导出 CSV
	csvExporter.Export("select * from xxx", &writer)

	oFetcher := NewOracleDataFetcher("mysql://127.0.0.1:1001")
	csvExporter.Fetcher(oFetcher)
	// 从 Oracle 数据源导出 CSV
	csvExporter.Export("select * from xxx", &writer)

	// 从 MySQL 数据源导出 JSON
	jsonExporter := NewJsonExporter(mFetcher)
	jsonExporter.Export("select * from xxx", &writer)
}
