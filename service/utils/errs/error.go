package errs

/*
var (
	dbErrorCode = map[int]error{
		1005: DBInsertFailed,    // CANNOT_CREATE_TABLE
		1006: DBInsertFailed,    // CANNOT_CREATE_DATABASE
		1007: DBInsertFailed,    // DATABASE_CREATE_EXISTS
		1040: DBOperationFailed, // TOO_MANY_CONNS
		1045: DBAccessDenied,    // ACCESS_DENIED
		1051: DBOperationFailed, // UNKNOWN_TABLE
		1062: DBInsertDuplicate, //DUPLICATE_ENTRY
	}
)

func ConvertDB(err error) error {
	if err == sql.ErrNoRows {
		return DBNoRows
	}

	if e, ok := parseDBError(err); ok {
		return e
	}

	return err
}

func parseDBError(err error) (error, bool) {
	s := strings.TrimSpace(err.Error())
	data := strings.Split(s, ":")
	if len(data) == 0 {
		return nil, false
	}

	numStr := strings.ToLower(data[0])
	numStr = strings.Replace(numStr, "error", "", -1)
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, false
	}

	e, ok := dbErrorCode[num]
	if !ok {
		return nil, false
	}

	return e, true
}
*/
