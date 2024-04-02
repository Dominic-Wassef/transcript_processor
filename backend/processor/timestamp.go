package processor

import "time"

// convertMsToRFC3339 converts milliseconds since epoch to an RFC3339 formatted timestamp.
func convertMsToRFC3339(ms int64) string {
	return time.Unix(0, ms*int64(time.Millisecond)).Format(time.RFC3339)
}
