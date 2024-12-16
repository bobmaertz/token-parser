package version

import "fmt"

var (
	Version string
	Build   string
)

func Print([]string) {
	fmt.Printf("Version: %s, Build %s", Version, Build)
}
