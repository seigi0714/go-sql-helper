package sqlmodel

type user struct{}

func NewUser() *user {
	return &user{}
}

func (u *user) Table() string {
	return "user"
}

func (u *user) PrimaryKey() string {
	return "id"
}

func (u *user) FieldsDef() []FieldDefinition {
	return []FieldDefinition{
		{"id", true, "user.id"},
		{"name", true, "user.name"},
		{"age", true, "user.age"},
	}
}

func (u *user) JoinTablesDef() []JoinDefinition {
	return []JoinDefinition{
		{"todo", "todo.userId = user.id"},
	}
}
