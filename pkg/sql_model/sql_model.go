package sqlmodel

type SqlModel interface {
	Table() string
	PrimaryKey() string
	FieldsDef() []FieldDefinition
	JoinTablesDef() []JoinDefinition
}

type FieldDefinition struct {
	Alias     string
	IsDefault bool
	Sql       string
}

type JoinDefinition struct {
	Alias string
	Sql   string
}
