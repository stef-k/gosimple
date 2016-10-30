package limiter

import (
	"time"
	"github.com/muesli/cache2go"
)

// InspectLoginAttempt records and checks if a login limit has been reached
// Returns true if a limit has been reached, false otherwise and the remaining attetmps
func RecordLoginAttempt(username string, incomingIP string) (bool, int, time.Time) {
	loginCache := cache2go.Cache("loginAttempt")
	if !loginCache.Exists(incomingIP) {
		var entry loginAttempt
		entry.FailedAttempts = loginAttempts
		entry.IncomingIP = incomingIP
		entry.Username = username
		entry.Timestamp = time.Now()
		loginCache.Add(incomingIP, time.Duration(loginLockMinutes) * time.Minute, &entry)
		return false, loginAttempts, time.Now()
	} else {
		limitReached, remainingAttempts, timestamp := checkLoginLimitReached(incomingIP)
		return limitReached, remainingAttempts, timestamp
	}
}

// checkLoginLimitReached checks if the login limit has been reached and returns true, false otherwise
func checkLoginLimitReached(incomingIP string) (bool, int, time.Time) {
	loginCache := cache2go.Cache("loginAttempt")
	if entry, ok := loginCache.Value(incomingIP); ok == nil {
		if entry.Data().(*loginAttempt).FailedAttempts == 0 {
			return true, entry.Data().(*loginAttempt).FailedAttempts, entry.Data().(*loginAttempt).Timestamp
		} else {
			return false, decreaseAttempt(incomingIP), entry.Data().(*loginAttempt).Timestamp
		}
	}
	return false, -1, time.Now()
}

// increaseAttempt increases the failed login attempts by 1
func decreaseAttempt(incomingIP string) int {
	loginCache := cache2go.Cache("loginAttempt")

	if entry, ok := loginCache.Value(incomingIP); ok == nil {
		entry.Data().(*loginAttempt).FailedAttempts -= 1
		entry.Data().(*loginAttempt).Timestamp = time.Now()
		return entry.Data().(*loginAttempt).FailedAttempts
	} else {
		return -1
	}
}
