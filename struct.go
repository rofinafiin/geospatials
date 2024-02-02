package geospatials

import "go.mongodb.org/mongo-driver/bson"

type GeoBorder struct {
	Type        string        `bson:"type"`
	Coordinates [][][]float64 `bson:"coordinates"`
}

type LocationData struct {
	ID          string    `bson:"_id"`
	Province    string    `bson:"province"`
	District    string    `bson:"district"`
	SubDistrict string    `bson:"sub_district"`
	Village     string    `bson:"village"`
	Border      GeoBorder `bson:"border"`
}

type RequestGeoIntersects struct {
	Coordinates [][][]float64 `bson:"coordinates" json:"coordinates"`
}

type Nearspherereq struct {
	Radius      int       `bson:"radius" json:"radius"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type GeoBoxReq struct {
	LowerLeft  []float64 `json:"lowerLeft" bson:"lowerLeft"`
	LowerRight []float64 `json:"lowerRight" bson:"lowerRight"`
}

type Geometryreq struct {
	Geometry bson.M `json:"geometry" bson:"geometry"`
}
