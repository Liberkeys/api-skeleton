package db

// Adaptor generic store
type Adaptor interface {
	ExecuteQuery(query string) error // Generic approach
}
