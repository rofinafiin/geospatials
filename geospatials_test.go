package geospatials

import (
	"fmt"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
)

func MongoCreateConnection(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func TestGeoIntersectQuery(t *testing.T) {
	client := MongoCreateConnection("mongodb+srv://rofinafiin:aXz4RdVqUVIQcqa1@rofinafiinsdata.9fyvx4r.mongodb.net/test", "GIS")
	polygonCoordinates := []float64{
		97.4471, 2.9622,
	}
	results, err := GeoMinDistanceQuery(client, polygonCoordinates, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Documents intersecting with the specified polygon:")
	for _, locationData := range results {
		fmt.Printf("%+v\n", locationData)
	}
}
