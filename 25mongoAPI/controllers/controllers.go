package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrExplorist/mongoAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "" // Here your mongo db connection string

const dbName = "netflix"
const collName = "watchlist"

 var collection *mongo.Collection

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	// collection instance is created to perform CRUD operations on the collection created.
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Collection instance created!")
}

// Mongodb helper functions -file


func createOneMovie(movie model.Netflix) (*mongo.InsertOneResult, error) {
	// insert one movie into database and return the inserted ID and error if any error occurs during insertion. 
	insertResult, err := collection.InsertOne(context.Background(), movie) // insert result is a struct type that contains the ID of the inserted document. 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult, nil
}

// update a movie in database
func updateOneMovie(movieId string){
	id , _ := primitive.ObjectIDFromHex(movieId) // convert to string 

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}


	result, err :=collection.UpdateOne(context.Background() , filter , update)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Updated a single document: ", result.ModifiedCount)}

// delete a movie in database
func deleteOneMovie(movieId string){
	id , _ := primitive.ObjectIDFromHex(movieId) // convert to string
	filter := bson.M{"_id": id}
	deleteCount , err := collection.DeleteOne(context.Background() , filter)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

// delete all movies in database

func deleteAllMovies() int64{
  deleteResult , err := collection.DeleteMany(context.Background() , bson.D{{}} , nil)



  if err != nil {
	log.Fatal(err)
  }

  fmt.Println("Movie deleted with delete count: ", deleteResult.DeletedCount)


  return deleteResult.DeletedCount

}

// get all movies from database


func getAllMovies() []primitive.M{

	//collection.Find method is used to query the MongoDB collection. In this case, it finds all documents in the collection by passing an empty bson.D{} filter. 
	cur , err := collection.Find(context.Background() , bson.D{{}})

	if err != nil{
		log.Fatal(err)
	}


	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie) //The & operator is used to pass a pointer to movie so that it can be modified by the Decode method.
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie) // movies slice is appended with the movie document using the append function.

	}
	defer cur.Close(context.Background())
	return movies
}

// Actual Controllers /  handlers for HTTP requests


func ServeHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to the home page</h1>"))

}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	createOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


// The DeleteAMovie function is an HTTP DELETE request handler. It sets the necessary headers for the response, extracts the movie ID from the route parameters, calls the deleteOneMovie function to delete the movie based on the ID, and then encodes the ID as JSON in the response. This function is intended to be used as a controller to handle DELETE requests for deleting movies in an API.
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"]) //  encodes the movie ID as JSON and writes it to the response writer (w). This provides a response to the client with the deleted movie's ID.
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}

