package src

type ActiveRecord struct {
	isNew bool
	table string
}

func (ar *ActiveRecord) tableName(table string) {
	ar.table = table
}
