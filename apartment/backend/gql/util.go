package gql

import "time"

func getNow(args map[string]interface{}) (string, error) {
	var now string
	if nowArg, ok := args["now"]; ok {
		nowArgStr := nowArg.(string)
		time, err := time.Parse(DateFormat, nowArgStr)
		if err != nil {
			return now, err
		}
		now = time.Format(DateFormat)
	} else {
		now = time.Now().Format(DateFormat)
	}
	return now, nil
}
