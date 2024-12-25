package main

import (
	"back-end/models"
	"back-end/services/get"
	"back-end/settings"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func enableSettings(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

func videos(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w)
	switch r.Method {
	case http.MethodGet:
		response := models.ServiceResponse{
			IsSuccessful: false,
			ErrorMessage: "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		videoID := pathParts[3]
		if videoID == "" {
			response.ErrorMessage = "No video id provided"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video, response := get.GetVideoByID(videoID)
		if !response.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(video)
		break
	case http.MethodPost:
		response := models.ServiceResponse{
			IsSuccessful: false,
			ErrorMessage: "",
		}
		var video models.Video
		err := json.NewDecoder(r.Body).Decode(&video)
		fmt.Println(video)
		if err != nil {
			response.ErrorMessage = "Video was not properly decoded" + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		break
	case http.MethodDelete:
		break
	case http.MethodPatch:
		break
	case http.MethodOptions:
		break
	}
}

func movies(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w)
	switch r.Method {
	case http.MethodGet:
		videosResponse, getResponse := get.GetVideosByType("movie")
		if !getResponse.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		json.NewEncoder(w).Encode(videosResponse)
		break
	}
}

func shows(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w)
	switch r.Method {
	case http.MethodGet:
		getResponse := models.ServiceResponse{
			IsSuccessful: false,
			ErrorMessage: "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		showID := pathParts[4]
		if showID == "" {
			getResponse.ErrorMessage = "No show id provided"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		videos, getResponse := get.GetVideosByShow(showID)
		if !getResponse.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		json.NewEncoder(w).Encode(videos)
		break
	}
}

func main() {
	http.HandleFunc(settings.VIDEOS_PATH, videos)
	http.HandleFunc(settings.MOVIES_PATH, movies)
	http.HandleFunc(settings.SHOWS_PATH, shows)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
