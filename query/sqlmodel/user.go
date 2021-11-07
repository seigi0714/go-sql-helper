package sqlmodel

type user struct{}

func NewUser() SqlModel {
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
		{"id", true, "user.id", "user"},
		{"name", true, "user.name", "user"},
		{"age", true, "user.age", "user"},
		{"todo_id", false, "todo.id", "todo"},
		{"post_id", false, "post.id", "post"},
	}
}

func (u *user) JoinTablesDef() []JoinDefinition {
	return []JoinDefinition{
		{"todo", "INNER JOIN todo ON todo.user_id = user.id"},
		{"post", "INNER JOIN post ON post.user_id = user.id"},
	}
}
