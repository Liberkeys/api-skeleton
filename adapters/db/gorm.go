package db

import "fmt"

// GormAdaptor ....
type GormAdaptor struct {
	Hostname   string
	Port       int
	GormClient string
}

// NewGormAdaptor ...
func NewGormAdaptor() *GormAdaptor {
	// Using Constructor force having all dependencies
	return &GormAdaptor{
		Hostname:   "localhost",
		Port:       3000,
		GormClient: "some client gorm instance...",
	}
}

// ExecuteQuery ...
func (db *GormAdaptor) ExecuteQuery(query string) error {
	fmt.Println("ADAPTER db.Execute :" + query)
	return nil
}
