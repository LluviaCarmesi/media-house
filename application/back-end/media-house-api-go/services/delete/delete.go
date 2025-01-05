package delete

import (
	"back-end/models"
	"back-end/services"
	"back-end/services/get"
	"back-end/settings"
	"context"
	"strconv"

	"github.com/google/uuid"
)

func DeleteVideo(videoID uuid.UUID) models.ServiceResponse {
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
