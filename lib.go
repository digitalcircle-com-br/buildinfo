package buildinfo

import (
	"fmt"
	"os"
)

var Ver string
var BuildDate string
var BuildHost string
var BuildUser string
var Product string

func String() string {

	extraInfo := ""
	for _, v := range []string{"buildinfo.yaml", "buildinfo.json", "buildinfo.txt", "buildinfo"} {
		_, err := os.Stat(v)
		if err != nil {
			bs, err := os.ReadFile(v)
			if err != nil {
				extraInfo = string(bs)
				break
			}
		}
	}

	return fmt.Sprintf(
		`%s: 
  - Ver: %s
  - Built at: %s
  - Built in Host: %s
  - Built by: %s
  %s`, Product, Ver, BuildDate, BuildHost, BuildUser, extraInfo)

}
