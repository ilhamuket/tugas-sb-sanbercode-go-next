package utils

import (
	"fmt"
	"time"
)

// ParseTime converts a database time value to a time.Time type.
func ParseTime(value interface{}) (time.Time, error) {
	// Database value might be in different formats (e.g., []uint8, string)
	switch v := value.(type) {
	case time.Time:
		return v, nil
	case []byte:
		return time.Parse("2006-01-02 15:04:05", string(v))
	case string:
		return time.Parse("2006-01-02 15:04:05", v)
	default:
		return time.Time{}, fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *time.Time", value)
	}
}
