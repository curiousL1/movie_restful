package response

import (
	"encoding/json"
	"net/http"

	"restful/models"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var dao = models.Movies{}

func httpResponseJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func FindMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := dao.FindMovieById(id)
	if err != nil {
		httpResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResponseJSON(w, http.StatusOK, result)
}

func FindMovieByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	result, err := dao.FindMovieByName(name)
	if err != nil {
		httpResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResponseJSON(w, http.StatusOK, result)
}

func FindAllMovies(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movies []models.Movies
	movies, err := dao.FindAllMovies()
	if err != nil {
		httpResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResponseJSON(w, http.StatusOK, movies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movies
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		httpResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	movie.Id = bson.NewObjectId()
	if err := dao.InsertMovie(movie); err != nil {
		httpResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResponseJSON(w, http.StatusCreated, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var params models.Movies
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		httpResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := dao.UpdateMovie(params); err != nil {
		httpResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResponseJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := dao.RemoveMovieById(id); err != nil {
		httpResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResponseJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
