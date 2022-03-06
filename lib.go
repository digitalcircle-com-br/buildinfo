package buildinfo

import "fmt"

var Ver string
var BuildDate string
var BuildHost string
var BuildUser string
var Product string

func String() string {
	return fmt.Sprintf(`Build Info:
		%s
		Ver: %s
		Built at: %s
		Built in Host: %s
		Built by: %s
	`, Product, Ver, BuildDate, BuildHost, BuildUser)
}
