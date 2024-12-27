package main

import (
	"back-end/models"
	"back-end/services/get"
	"back-end/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

		// parsing all form data
		r.ParseMultipartForm(50 << 20) // 50MB limit
		videoFile, videoFileHeader, err := r.FormFile("VideoFile")
		if err != nil {
			response.ErrorMessage = "Video File is not proper: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer videoFile.Close()
		var videoFileBuffer bytes.Buffer
		_, err = io.Copy(&videoFileBuffer, videoFile)
		if err != nil {
			response.ErrorMessage = "Video File can't be read: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.VideoFile = videoFileBuffer.Bytes()
		video.VideoFileName = videoFileHeader.Filename

		previewFile, previewFileHeader, err := r.FormFile("PreviewFile")
		if err != nil {
			response.ErrorMessage = "Preview File is not proper: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer previewFile.Close()
		var previewFileBuffer bytes.Buffer
		_, err = io.Copy(&previewFileBuffer, previewFile)
		if err != nil {
			response.ErrorMessage = "Preview File can't be read: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.PreviewFile = previewFileBuffer.Bytes()
		video.PreviewFileName = previewFileHeader.Filename

		showIDInt, err := strconv.Atoi(r.FormValue("ShowID"))
		if err != nil {
			response.ErrorMessage = "Show ID is not an int: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.ID = ""
		video.Title = r.FormValue("Title")
		video.Type = r.FormValue("Type")
		video.Episode = r.FormValue("Episode")
		video.ShowID = showIDInt
		video.Duration = r.FormValue("Duration")
		video.Language = r.FormValue("Language")
		video.Tags = r.Form["Tags"]

		fmt.Println(video)
		json.NewEncoder(w).Encode(response)
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
