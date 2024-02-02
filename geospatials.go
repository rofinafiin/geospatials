package geospatials

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GeoIntersectQuery(client *mongo.Database, polygon [][][]float64) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")

	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": polygon,
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoWithinQuery(client *mongo.Database, polygon [][][]float64) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": polygon,
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoNearQuery(client *mongo.Database, polygon []float64, maxDistance, minDistance int) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": polygon,
				},
				"$maxDistance": maxDistance,
				"$minDistance": minDistance,
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoNearSphereQuery(client *mongo.Database, polygon []float64, radius int) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": polygon,
				},
				"$maxDistance": radius,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoBoxQuery(client *mongo.Database, lowerLeft, upperRight []float64) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"$box": [][]float64{
						lowerLeft,
						upperRight,
					},
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoCenterQuery(client *mongo.Database, center []float64, radius int) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")

	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{center, float64(radius) / 6371000},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoGeometryQuery(client *mongo.Database, geometry bson.M) ([]FullGeoJson, error) {
	// Actual implementation of the function
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": geometry,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
func GeoMaxDistanceQuery(client *mongo.Database, point []float64, maxDistance int) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry":    bson.M{"type": "Point", "coordinates": point},
				"$maxDistance": maxDistance,
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GeoMinDistanceQuery(client *mongo.Database, point []float64, minDistance int) ([]FullGeoJson, error) {
	collection := client.Collection("geogis")

	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry":    bson.M{"type": "Point", "coordinates": point},
				"$minDistance": minDistance,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []FullGeoJson
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
