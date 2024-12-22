package main

import (
	"back-end/models"
	"back-end/services/get"
	"back-end/settings"
	"encoding/json"
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
		getResponse := models.ServiceResponse{
			IsSuccessful: false,
			ErrorMessage: "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		videoID := pathParts[3]
		if videoID == "" {
			getResponse.ErrorMessage = "No video id provided"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		video, getResponse := get.GetVideoByID(videoID)
		if !getResponse.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		json.NewEncoder(w).Encode(video)
		break
	case http.MethodPost:
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
		videos, getResponse := get.GetVideosByType("movie")
		if !getResponse.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		json.NewEncoder(w).Encode(videos)
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
