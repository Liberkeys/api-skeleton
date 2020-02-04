package db

import "fmt"

// Driver ....
type Driver struct {
	uri string
}

// NewDriver ...
func NewDriver(host string, port string, user string, password string, dbname string) *Driver {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=disable",
		host, port, user, dbname, password)

	return &Driver{
		uri: uri,
	}
}
