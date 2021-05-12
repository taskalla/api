package paginate

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func SerializeCursor(index int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor:%d", index)))
}

func DeserializeCursor(cursor string) (int, error) {
	cursor_bytes, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, strings.NewReader(cursor)))
	if err != nil {
		return 0, errors.New("Invalid cursor")
	}
	cursor_string := string(cursor_bytes)

	split := strings.Split(cursor_string, ":")
	if len(split) != 2 {
		return 0, errors.New("Invalid cursor")
	}

	index, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, errors.New("Invalid cursor")
	}

	if index < 0 {
		return 0, errors.New("Invalid cursor")
	}

	return index, nil
}
