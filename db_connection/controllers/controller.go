package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pran2805/go-lang-learning/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URL string = "mongodb://127.0.0.1:27017/"
const dbName string = "go-learning"
const collectionName string = "watchlist"

// most important
var collection *mongo.Collection

// connect with mongodb
func init() {
	// client options
	clientOption := options.Client().ApplyURI(MONGODB_URL)

	// connection to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb Connection Success")

	collection = client.Database(dbName).Collection(collectionName)

	// connection instance
	fmt.Println("Connection Instance is ready")
}

type Data struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	MovieName string             `json:"movieName"`
	Watched   bool               `json:"watched"`
}

// Mongodb helpers - keep in seperate file
// insert 1 record
func insertMovie(movie models.Netflix) Data {
	insertData, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Movie data in db", insertData)
	id, ok := insertData.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("InsertedID is not a primitive.ObjectID")
	}

	data := Data{
		ID:        id,
		MovieName: movie.Movie,
		Watched:   movie.Watched,
	}
	return data
}

// update 1 record
func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	update := bson.M{"$set": bson.M{"watched": true}}

	fmt.Println(update)
	result, err := collection.UpdateByID(context.Background(), id, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated Result", result)
}

func deleteMovieById(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Result", result)
}

func deleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Result", result)

	return result.DeletedCount
}

// get all movies from database
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	fmt.Println(movies)

	return movies
}

// Actual Controller's
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	data := insertMovie(movie)
	json.NewEncoder(w).Encode(data)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteMovieById(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)

}
