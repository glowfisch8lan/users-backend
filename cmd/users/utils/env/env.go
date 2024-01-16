package env

import (
	"fmt"
	"os"
)

func Get(field string) any {

	field, exists := os.LookupEnv(field)
	if !exists {
		panic(fmt.Sprintf("%s not exist", field))
	}

	return field
}
