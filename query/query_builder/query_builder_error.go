package query

import "strconv"

const NOT_DEF_FUILD_CODE = 101
const NOT_DEF_JOIN_TABLE_CODE = 102
const INVALID_OPERATOR_CODE = 103

type BuildingError struct {
	code int
	msg  string
}

func (be *BuildingError) Error() string {
	return "failed to generate sql" + "\nCode :" + strconv.Itoa(be.code) + "\nMessage: " + be.msg
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

func InvalidOperatorError(oparator string) error {
	return &BuildingError{
		code: INVALID_OPERATOR_CODE,
		msg:  "Invalid Operator " + oparator,
	}
}
