package server

import (
	"fmt"
	"sync"
)

// Log 로그 기록
type Log struct {
	mu      sync.Mutex
	records []Record
}

// ErrOffsetNotFound Offset 을 찾을 수 없을 때 노출하는 에러메세지
var ErrOffsetNotFound = fmt.Errorf("offset not found")

// Record Log의 각 항목
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// NewLog 로그 생성
func NewLog() *Log {
	return &Log{}
}

// Append 기존 로그에 레코드 추가
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read 해당 오프셋에서 항목 가져오기
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
