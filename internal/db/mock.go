package db

import (
	"database/sql"
	"fmt"
)

type MockDB struct {}

type mockResult struct {}

func (m mockResult) LastInsertId() (int64, error) {return 1, nil}
func (m mockResult) RowsAffected() (int64, error) {return 1, nil}

func (m *MockDB) Connect(connectionString string) error {
	fmt.Println("[MockDB] Connect called with: ", connectionString)
	return nil
}

func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println("[MockDB] Query called with: ", query, args)
	return nil, nil
}


func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	fmt.Println("[MockDB] Exec called with: ", query, args)
	return mockResult{}, nil
}

func (m *MockDB) Close() error {
	fmt.Println("[MockDB] Close called")
	return nil
}


