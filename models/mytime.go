package models

import (
	"fmt"
	"time"
)

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02T15:04:05.999999"`, string(b))
	if err != nil {
		parsedTime, err = time.Parse(`"2006-01-02T15:04:05"`, string(b))
	}
	*t = MyTime(parsedTime)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02T15:04:05.999999"))), nil
}
