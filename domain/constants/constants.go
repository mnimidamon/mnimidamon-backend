package constants

import (
	"log"
)

// Replaceable application wide logger.
var Log func(format string, a ...interface{}) = log.Printf

