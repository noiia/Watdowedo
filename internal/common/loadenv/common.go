package loadenv

import "strconv"

func FormatDbFieldsToStruct(Host, Port, Name, User, Password, MaxRetry, RetryDelay string) (DbFields, error) {
	var dbMaxRetryInt int
	var dbRetryDelayInt int
	var err error

	if MaxRetry != "" {
		dbMaxRetryInt, err = strconv.Atoi(MaxRetry)
		if err != nil {
			return DbFields{}, err
		}
	}

	if RetryDelay != "" {
		dbRetryDelayInt, err = strconv.Atoi(RetryDelay)
		if err != nil {
			return DbFields{}, err
		}
	}
	return DbFields{Host, Port, Name, User, Password, dbMaxRetryInt, dbRetryDelayInt}, nil
}
