package dynamodb

type DB struct {}

func NewDB() IDynamoDB {
	return &DB{}
}

func (d *DB) CreateTable() error {
	return nil
}