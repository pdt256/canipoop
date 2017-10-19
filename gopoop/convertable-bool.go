package gopoop

import (
	"errors"
	"fmt"
)

type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
    asString := string(data)
    if asString == "1" || asString == "true" {
        *bit = true
    } else if asString == "0" || asString == "false" {
        *bit = false
    } else {
        return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
    }
    return nil
}
