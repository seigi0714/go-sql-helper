package sqlmodel

const LEFTJOIN = "LEFT JOIN"
const RIGTHJOIN = "RIGHT JOIN"
const INNERJOIN = "INNER JOIN"

type SqlModel interface {
	Table() string
	PrimaryKey() string
	FieldsDef() []FieldDefinition
	JoinTablesDef() []JoinDefinition
}

type FieldDefinition struct {
	Alias      string
	IsDefault  bool
	Sql        string
	TableAlias string
}

type JoinDefinition struct {
	Alias string
	Sql   string
}
