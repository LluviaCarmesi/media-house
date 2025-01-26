package main

import (
	"back-end/models"
	"back-end/services/delete"
	"back-end/services/get"
	"back-end/services/post"
	"back-end/settings"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
			Message:      "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		if pathParts[3] == "" {
			limitValue := r.URL.Query().Get("limit")
			offsetValue := r.URL.Query().Get("offset")
			videosResponse, serviceResponse := get.GetAllVideos(limitValue, offsetValue)
			if !serviceResponse.IsSuccessful {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(serviceResponse)
				return
			}
			json.NewEncoder(w).Encode(videosResponse)
			return
		}
		videoID, err := uuid.Parse(pathParts[3])
		if err != nil {
			response.Message = "id is not a uuid"
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
			Message:      "",
		}
		var video models.Video
		var videoChunks models.VideoChunks

		// parsing all form data
		r.ParseMultipartForm(50 << 20) // 50MB limit
		videoFile, videoFileHeader, err := r.FormFile("VideoFile")
		if err != nil {
			response.Message = "Video File is not proper: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer videoFile.Close()
		var videoFileBuffer bytes.Buffer
		_, err = io.Copy(&videoFileBuffer, videoFile)
		if err != nil {
			response.Message = "Video File can't be read: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.VideoFile = videoFileBuffer.Bytes()
		video.VideoFileName = videoFileHeader.Filename

		previewFile, previewFileHeader, err := r.FormFile("PreviewFile")
		if err != nil {
			response.Message = "Preview File is not proper: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer previewFile.Close()
		var previewFileBuffer bytes.Buffer
		_, err = io.Copy(&previewFileBuffer, previewFile)
		if err != nil {
			response.Message = "Preview File can't be read: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.PreviewFile = previewFileBuffer.Bytes()
		video.PreviewFileName = previewFileHeader.Filename

		showIDInt, err := strconv.Atoi(r.FormValue("ShowID"))
		if err != nil {
			response.Message = "Show ID is not an int: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		video.ID = uuid.New()
		video.Title = r.FormValue("Title")
		video.Type = r.FormValue("Type")
		video.Episode = r.FormValue("Episode")
		video.Duration = r.FormValue("Duration")
		video.Language = r.FormValue("Language")

		tagsString := r.FormValue("Tags")

		err = json.Unmarshal([]byte(tagsString), &video.Tags)
		if err != nil {
			response.Message = "Couldn't parse tags: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		if showIDInt != 0 {
			video.ShowID = &showIDInt
		}

		videoFileChunksNumberInt, err := strconv.Atoi(r.FormValue("VideoFileChunkNumber"))
		if err != nil {
			response.Message = "VideoChunkNumber is not an int: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		videoFileTotalChunksInt, err := strconv.Atoi(r.FormValue("VideoFileTotalChunks"))
		if err != nil {
			response.Message = "VideoTotalChunks is not an int: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		videoChunks.VideoFileChunkNumber = videoFileChunksNumberInt
		videoChunks.VideoFileTotalChunks = videoFileTotalChunksInt

		addVideoResponse := post.AddVideoFiles(video, videoChunks)
		if !addVideoResponse.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(addVideoResponse)
			return
		}
		json.NewEncoder(w).Encode(response)
		break
	case http.MethodDelete:
		response := models.ServiceResponse{
			IsSuccessful: false,
			Message:      "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		videoID, err := uuid.Parse(pathParts[3])
		if err != nil {
			response.Message = "id is not a uuid"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		if videoID.String() == "" {
			response.Message = "No id provided"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		response = delete.DeleteVideoFiles(videoID)
		if !response.IsSuccessful {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(response)
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
		limitValue := r.URL.Query().Get("limit")
		offsetValue := r.URL.Query().Get("offset")
		videosResponse, getResponse := get.GetVideosByType("movie", limitValue, offsetValue)
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
			Message:      "",
		}
		pathParts := strings.Split(r.URL.Path, "/")
		showID := pathParts[4]
		if showID == "" {
			getResponse.Message = "No show id provided"
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

func search(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w)
	switch r.Method {
	case http.MethodGet:
		getResponse := models.ServiceResponse{
			IsSuccessful: false,
			Message:      "",
		}
		limitValue := r.URL.Query().Get("limit")
		offsetValue := r.URL.Query().Get("offset")
		pathParts := strings.Split(r.URL.Path, "/")
		searchTerm := pathParts[4]
		if searchTerm == "" {
			getResponse.Message = "No search term provided"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(getResponse)
			return
		}
		videos, getResponse := get.GetVideosBySearch(searchTerm, limitValue, offsetValue)
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
	http.HandleFunc(settings.VIDEOS_SEARCH_PATH, search)

	log.Fatal(http.ListenAndServe(settings.SERVE_PORT, nil))
}
