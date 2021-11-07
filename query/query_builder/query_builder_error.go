package query

import "strconv"

const NOT_DEF_FUILD_CODE = 101
const NOT_DEF_JOIN_TABLE_CODE = 102

type BuildingError struct {
	code int
	msg  string
}

func (be *BuildingError) Error() string {
	return "failed to generate" + "\nCode :" + strconv.Itoa(be.code) + "\nMessage: " + be.msg
}

func NotDefFieldError(field string) error {
	return &BuildingError{
		code: NOT_DEF_FUILD_CODE,
		msg:  "Not Definition Field " + field,
	}
}

func NotDefJoinTableError(table string) error {
	return &BuildingError{
		code: NOT_DEF_JOIN_TABLE_CODE,
		msg:  "Not Definition Join Table " + table,
	}
}
