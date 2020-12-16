package models

import (
	//"time"
	//"fmt"
	"github.com/globalsign/mgo/bson"
)


type Movies struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Time 		string     	  `bson:"time" json:"time"`
	Description string        `bson:"description" json:"description"`
}

const (
	database     = "Movies"
	dbcollection = "MovieData"
)

func (m *Movies) InsertMovie(movie Movies) error {
	return Insert(database, dbcollection, movie)
}

func (m *Movies) FindAllMovies() ([]Movies, error) {
	var result []Movies
	err := FindAll(database, dbcollection, nil, nil, &result)

	//fmt.Printf("FindAllMovies:%v \n", &result)
	return result, err
}

func (m *Movies) FindMovieById(id string) (Movies, error) {
	var result Movies
	err := Find(database, dbcollection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)

	//fmt.Printf("FindMovieById:%v \n", &result)
	return result, err
}

func (m *Movies) FindMovieByName(name string) (Movies, error) {
	var result Movies
	err := Find(database, dbcollection, bson.M{"name": name}, nil, &result)
	return result, err
}

func (m *Movies) UpdateMovie(movie Movies) error {
	return Update(database, dbcollection, bson.M{"_id": movie.Id}, movie)
}

func (m *Movies) RemoveMovieById(id string) error {
	return Remove(database, dbcollection, bson.M{"_id": bson.ObjectIdHex(id)})
}
