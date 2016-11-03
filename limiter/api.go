package limiter

import (
	"time"
	"github.com/muesli/cache2go"
)

// ApiLimitReached checks if the login limit is reached for a specific IP
// Return true if login attempts reached the limit, false otherwise.
func ApiLimitReached(incomingIP string) bool {
	apiCache := cache2go.Cache("incrementRequest")
	if entry, ok := apiCache.Value(incomingIP);ok == nil {
		return entry.Data().(*apiRequest).CurrentRequests > apiRequestsLimit
	} else {
		return false
	}
}

// ApiRecordRequest records a new API request
// Returns the current requests and the request limit
func ApiRecordRequest(incomingIP string) (currentRequests, requestsLimit int) {
	apiCache := cache2go.Cache("incrementRequest")
	currentRequests = 1
	requestsLimit = apiRequestsLimit
	if !apiCache.Exists(incomingIP) {
		var entry apiRequest
		entry.CurrentRequests = currentRequests
		if limitingTimeUnit == "min" {
			apiCache.Add(incomingIP, time.Duration(perTime) * time.Minute, &entry)
		} else {
			apiCache.Add(incomingIP, time.Duration(perTime) * time.Second, &entry)
		}
		return currentRequests, requestsLimit
	} else {
		if entry, ok := apiCache.Value(incomingIP); ok == nil {
			if entry.Data().(*apiRequest).CurrentRequests <= apiRequestsLimit {
				entry.Data().(*apiRequest).CurrentRequests += 1
			}
			currentRequests = entry.Data().(*apiRequest).CurrentRequests
		}
		return currentRequests, requestsLimit
	}
}
