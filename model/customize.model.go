package model

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Customize struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	KartIndex   int64       `json:"kartIndex" bson:"kartIndex"`
	OutfitIndex int64       `json:"outfitIndex" bson:"outfitIndex"`
}
