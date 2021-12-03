package datetime

import "time"

// FormatDate returns time formatted as RFC3339 (2021-11-10T17:41:48+09:00)
func FormatDate(t time.Time) string {
	return t.Format(time.RFC3339)
}
