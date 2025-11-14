package utilities

import (
	"back-end/models"
	"back-end/settings"
	"fmt"
	"os"
	"strings"
)

func CheckVideoModel(video models.Video) models.ModelCheckResponse {
	response := models.ModelCheckResponse{
		IsValid:  false,
		Response: "",
	}

	if video.VideoFile == nil {
		response.Response = "Video File is required"
		return response
	}
	if video.PreviewFile == nil {
		response.Response = "Preview File is required"
		return response
	}
	if video.Title == "" {
		response.Response = "Title is required"
		return response
	}
	if video.Duration == "" {
		response.Response = "Duration is required"
		return response
	}
	if video.Language == "" {
		response.Response = "Language is required"
		return response
	}
	if video.Type == "" {
		response.Response = "Type is required"
		return response
	}

	if video.Type == settings.SHOWS_TYPE {
		if video.Episode == "" {
			response.Response = "Episode is required"
			return response
		}
		if *video.ShowID == 0 || video.ShowID == nil {
			response.Response = "Show is required"
			return response
		}
	}

	response.IsValid = true
	return response
}

func IsSearchTermInVideo(searchTerm string, video models.Video) bool {
	doesVideoHaveSearchTerm := false
	if strings.Contains(strings.ToLower(video.Title), searchTerm) {
		doesVideoHaveSearchTerm = true
	}

	for i := 0; i < len(video.Tags); i++ {
		currentTag := strings.ToLower(video.Tags[i])
		if strings.Contains(currentTag, searchTerm) {
			doesVideoHaveSearchTerm = true
		}
	}

	return doesVideoHaveSearchTerm
}

func IsValidImageFile(path string) bool {
	f, err := os.Open(settings.MEDIA_DIRECTORY + path)
	fmt.Println(f)
	fmt.Println(err)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}
