package promotions

var PromotionCollection = "promotions"

type Promotion struct {
	Id              string `bson:"_id"`
	Price           float64
	Expiration_date string
}
