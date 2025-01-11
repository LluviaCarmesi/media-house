package delete

import (
	"back-end/models"
	"back-end/services"
	"back-end/services/get"
	"back-end/settings"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func deleteVideo(videoID uuid.UUID) models.ServiceResponse {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	videoTags, videoTagsResponse := get.GetVideoTagsByVideoID(videoID)
	if !videoTagsResponse.IsSuccessful {
		response = videoTagsResponse
		return response
	}

	for i := 0; i < len(videoTags); i++ {
		deleteVideoResponse := DeleteVideoTag(strconv.Itoa(videoTags[i].ID))
		if !deleteVideoResponse.IsSuccessful {
			response = deleteVideoResponse
			return response
		}
	}

	query := settings.DELETE_VIDEO_QUERY
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		videoID,
	)
	if err != nil {
		response.Message = "Video couldn't be deleted: " + err.Error()
		return response
	}

	response.IsSuccessful = true
	response.Message = "Video was deleted successfully"
	return response
}

func DeleteVideoTag(id string) models.ServiceResponse {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	query := settings.DELETE_VIDEO_TAG_QUERY
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id,
	)
	if err != nil {
		response.Message = "Video couldn't be deleted: " + err.Error()
		return response
	}

	response.IsSuccessful = true
	response.Message = "Video was deleted successfully"
	return response
}

func DeleteVideoFiles(videoID uuid.UUID) models.ServiceResponse {
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	video, videoResponse := get.GetVideoByID(videoID)
	if !videoResponse.IsSuccessful {
		response = videoResponse
		return response
	}
	previewPathSplit := strings.Split(video.PreviewPath, "/")
	videoPathSplit := strings.Split(video.VideoPath, "/")
	fmt.Println(previewPathSplit)
	fmt.Println(videoPathSplit)
	fmt.Println(len(previewPathSplit))
	if len(previewPathSplit) > 1 {
		previewFileDestinationPath := settings.PREVIEW_FILES_DIRECTORY + previewPathSplit[1]
		_, err := os.Stat(previewFileDestinationPath)
		if err != nil {
			response.Message = "Preview File doesn't exist: " + err.Error()
			return response
		}
		err = os.Remove(previewFileDestinationPath)
		if err != nil {
			response.Message = "Preview File couldn't be removed: " + err.Error()
			return response
		}
	}

	if len(videoPathSplit) > 1 {
		videoFileDestinationPath := settings.VIDEO_FILES_DIRECTORY + videoPathSplit[1]
		_, err := os.Stat(videoFileDestinationPath)
		if err != nil {
			response.Message = "Video doesn't exist: " + err.Error()
			return response
		}
		err = os.Remove(videoFileDestinationPath)
		if err != nil {
			response.Message = "Video File couldn't be removed: " + err.Error()
			return response
		}
	}

	deleteVideoResponse := deleteVideo(videoID)
	if !deleteVideoResponse.IsSuccessful {
		response = deleteVideoResponse
		return response
	}

	response.IsSuccessful = true
	response.Message = "Video was successfully removed"
	return response
}
