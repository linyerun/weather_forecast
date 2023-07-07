package db

import (
	"fmt"
)

func AutoTables(objs ...any) {
	for _, obj := range objs {
		if err := db.AutoMigrate(obj); err != nil {
			fmt.Println(obj)
			continue
		}
	}
}
