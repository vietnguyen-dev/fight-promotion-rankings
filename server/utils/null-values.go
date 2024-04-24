package utils

import (
	"database/sql"
	"time"
)

// Helper function to convert sql.NullString to a plain string
func NullStringToString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return "" // Return an empty string if the value is invalid (NULL)
}


func NullInt64ToInt(ni sql.NullInt64) int {
    if ni.Valid {
        return int(ni.Int64) // Convert the int64 to int
    }
    return 0 // Return a default value if the value is invalid (NULL)
}

func NullTimeToTime(nt sql.NullTime) time.Time {
    if nt.Valid {
        return nt.Time // Return the date if valid
    }
    return time.Time{} // Return zero value of time.Time if NULL
}

func NullBoolToBool(nb sql.NullBool) bool {
    if nb.Valid {
        return nb.Bool // Return the boolean if valid
    }
    return false // Return false if NULL
}
