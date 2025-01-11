package post

import (
	"back-end/models"
	"back-end/services"
	"back-end/settings"
	"back-end/utilities"
	"context"
	"io"
	"log"
	"os"
	"strconv"
)

func addVideo(video models.Video) models.ServiceResponse {
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()

	query := settings.INSERT_VIDEO_QUERY
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		video.ID,
		video.Title,
		video.Type,
		video.Episode,
		video.ShowID,
		video.Duration,
		video.Language,
		video.PreviewPath,
		video.VideoPath,
	)
	if err != nil {
		response.Message = "Video details couldn't be added: " + err.Error()
		return response
	}

	log.Printf("Inserted id for videos: %d", video.ID)

	for i := 0; i < len(video.Tags); i++ {
		if video.Tags[i] != "" {
			addVideoTagResponse := AddVideoTag(models.VideoTag{
				Title:   video.Tags[i],
				VideoID: video.ID,
			})
			if !addVideoTagResponse.IsSuccessful {
				response.Message = addVideoTagResponse.Message
				return response
			}
		}
	}

	response.IsSuccessful = true
	response.Message = "Video was added successfully"
	return response
}

func AddVideoFiles(video models.Video, videoChunks models.VideoChunks) models.ServiceResponse {
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	videoFileDestinationPath := settings.VIDEO_FILES_DIRECTORY + video.VideoFileName

	_, err := os.Stat(videoFileDestinationPath)
	if err == nil {
		response.Message = "Video already exists"
		return response
	}

	videoModelCheck := utilities.CheckVideoModel(video)
	if !videoModelCheck.IsValid {
		response.Message = videoModelCheck.Response
		return response
	}

	err = os.WriteFile(settings.VIDEO_FILES_TEMP_DIRECTORY+
		video.VideoFileName+
		"_"+
		strconv.Itoa(videoChunks.VideoFileChunkNumber),
		video.VideoFile,
		0755,
	)
	if err != nil {
		response.Message = "Couldn't add video to directory: " + err.Error()
		return response
	}

	if videoChunks.VideoFileChunkNumber != (videoChunks.VideoFileTotalChunks - 1) {
		response.IsSuccessful = true
		response.Message = "Chunk was added"
		return response
	}

	videoFileDestination, err := os.Create(videoFileDestinationPath)
	if err != nil {
		response.Message = "Video file couldn't be created: " + err.Error()
		return response
	}
	previewFileDestinationPath := settings.PREVIEW_FILES_DIRECTORY + video.PreviewFileName
	err = os.WriteFile(previewFileDestinationPath,
		video.PreviewFile,
		0755,
	)
	if err != nil {
		response.Message = "Couldn't add preview to directory: " + err.Error()
		return response
	}

	for i := 0; i < videoChunks.VideoFileTotalChunks; i++ {
		chunkPath := settings.VIDEO_FILES_TEMP_DIRECTORY + video.VideoFileName + "_" + strconv.Itoa(i)
		currentVideoChunk, err := os.Open(chunkPath)
		if err != nil {
			response.Message = "Chunk " + strconv.Itoa(i) + "couldn't be opened: " + err.Error()
			return response
		}
		_, err = io.Copy(videoFileDestination, currentVideoChunk)
		if err != nil {
			response.Message = "Chunk " + strconv.Itoa(i) + "couldn't be copied: " + err.Error()
			return response
		}
		err = os.Remove(chunkPath)
		if err != nil {
			response.Message = "Chunk " + strconv.Itoa(i) + "couldn't be removed: " + err.Error()
			return response
		}
		currentVideoChunk.Close()
	}

	video.VideoPath = settings.VIDEO_FILES_SERVE_PATH + video.VideoFileName
	video.PreviewPath = settings.PREVIEW_FILES_SERVE_PATH + video.PreviewFileName
	addVideoResponse := addVideo(video)
	if !addVideoResponse.IsSuccessful {
		response.Message = addVideoResponse.Message
		return response
	}

	response.IsSuccessful = true
	response.Message = "Video was successfully added"
	return response
}

func AddVideoTag(videoTag models.VideoTag) models.ServiceResponse {
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()

	query := settings.INSERT_VIDEO_TAG_QUERY
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		videoTag.Title,
		videoTag.VideoID,
	)
	if err != nil {
		response.Message = "Unable to insert video tag: " + err.Error()
		return response
	}

	insertedID, err := results.LastInsertId()
	if err != nil {
		log.Fatalf("Impossible to retrieve last inserted id for video_tags: %s", err.Error())
	}
	log.Printf("Inserted id for video_tags: %d", insertedID)

	response.IsSuccessful = true
	response.Message = "Video Tag was added successfully"
	return response
}

func AddShow(show models.Show) models.ServiceResponse {
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()

	query := settings.INSERT_SHOW_QUERY
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		show.Title,
		show.PreviewPath,
	)
	if err != nil {
		response.Message = "Unable to insert show: " + err.Error()
		return response
	}

	insertedID, err := results.LastInsertId()
	if err != nil {
		log.Fatalf("Impossible to retrieve last inserted id for shows: %s", err.Error())
	}
	log.Printf("Inserted id for shows: %d", insertedID)

	return response
}
