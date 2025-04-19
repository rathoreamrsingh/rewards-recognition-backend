package model

type Points struct {
	ID               int `json:"_id" bson:"_id"`
	UserId           int `json:"user_id" bson:"user_id"`
	GivablePoints    int `json:"givable_points" bson:"givable_points"`
	RedeemablePoints int `json:"redeemable_points" bson:"redeemable_points"`
}
