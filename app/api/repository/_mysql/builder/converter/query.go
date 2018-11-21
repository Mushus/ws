package converter

// Query SQL文
type Query struct {
	sql  string
	args []interface{}
}
