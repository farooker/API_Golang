package model

type Avatar struct {
    Id 		    string  `json:"_id" bson:"_id"`
    ContactLens string  `json:"contactLens"`
    Emotes  	[]string  `json:"emotes"`
    VictoryPose string  `json:"victoryPose"`
    OwnerId  	string  `json:"ownerId"  bson:"ownerId"`
    Skin  		string  `json:"skin"`
    Eyebrows  	string  `json:"eyebrows"`
    Mouth  		string  `json:"mouth"`
    EyeShape 	string  `json:"eyeShape"`
}