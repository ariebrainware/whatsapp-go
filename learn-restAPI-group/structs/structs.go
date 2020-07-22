// learn rest api sprated into group (add config)

package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	firtName string
	lastName string
}
