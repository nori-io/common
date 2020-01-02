package storage

type (
	Storage interface {
		SQL() SQL
		KV(schema string) KeyValue
	}

	KeyValue interface {
		Get(key string) ([]byte, error)
		Set(key string, value []byte) error
		Delete(key string) error
	}

	SQL interface {
		Create(item interface{}) error
		Read(item interface{}) error
		Update(item interface{}) error
		Delete(item interface{}) error
		Find(out interface{}, where ...interface{}) error
		Migrate() Migrate
	}

	Migrate interface {
		HasTable(name string) bool
		CreateTable(models ...interface{}) error
		DropTable(models ...interface{}) error
		Migrate(models ...interface{})
	}

	Table interface {
		TableName() string
	}
)
