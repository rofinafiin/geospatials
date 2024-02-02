package geospatials

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type RequestGeonear struct {
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	MinDistance int       `bson:"minDistance" json:"minDistance"`
	MaxDistance int       `bson:"maxDistance" json:"maxDistance"`
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

type GeoPoint struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type GeoPolygon struct {
	Type        string        `json:"type" bson:"type"`
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
}

type GeoFeature struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Type       string             `json:"type" bson:"type"`
	Geometry   GeoPolygon         `json:"geometry" bson:"geometry"`
	Properties Name               `json:"properties" bson:"properties"`
}

type Name struct {
	Name string `bson:"name" json:"name"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

type FullGeoJson struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   Geometry           `json:"geometry" bson:"geometry"`
}

type Properties struct {
	Name string `json:"name" bson:"name"`
}
