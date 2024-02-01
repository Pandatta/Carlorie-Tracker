package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

typo Entry struct {
	ID			primitive.ObjectID `bson:"id"`
	Dish		*string 		'json:"dish"'
	Fat			*float64		'json:"fat"'
	Ingredients *string			'json:"ingredients"'
	proteins	*string		'json:"proteins"'
	fiber		*string			'json:"fiber"'
	carb		*string			'json:"carb"'
	Carlories	*string			'json:"carlories"'
	
}