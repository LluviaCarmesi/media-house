package get

import (
	"back-end/models"
	"back-end/services"
	"back-end/settings"
	"back-end/utilities"
	"strconv"

	"github.com/google/uuid"
)

func GetAllVideos(limitValue string, offsetValue string) (models.VideosResponse, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videos := []models.Video{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	numberOfVideos := 0
	allVideoTags, videoTagsResponse := GetVideoTags()
	if !videoTagsResponse.IsSuccessful {
		response.Message = "Error getting tags " + videoTagsResponse.Message
		return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		}, response
	}
	allShows, showsResponse := GetVideoTags()
	if !showsResponse.IsSuccessful {
		response.Message = "Error getting tags " + showsResponse.Message
		return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		}, response
	}

	query := settings.GET_ALL_VIDEOS_QUERY
	if limitValue != "" && offsetValue != "" {
		query += " LIMIT " + limitValue + " OFFSET " + offsetValue
	} else if limitValue != "" {
		query += " LIMIT " + limitValue

	} else if offsetValue != "" {
		query += " OFFSET " + offsetValue
	}

	results, err := dbConnection.Query(query)
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		}, response
	}

	for results.Next() {
		var video = models.Video{}

		err := results.Scan(
			&video.ID,
			&video.Title,
			&video.Type,
			&video.Episode,
			&video.ShowID,
			&video.Duration,
			&video.Language,
			&video.PreviewPath,
			&video.VideoPath,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}
		currentVideoTags := []string{}
		for i := 0; i < len(allVideoTags); i++ {
			if video.ID == allVideoTags[i].VideoID {
				currentVideoTags = append(currentVideoTags, allVideoTags[i].Title)
			}
		}
		currentVideoShow := ""
		if video.ShowID != nil {
			for i := 0; i < len(allShows); i++ {
				if *video.ShowID == allShows[i].ID {
					currentVideoShow = allShows[i].Title
				}
			}
		}
		if !utilities.IsValidImageFile(video.PreviewPath) {
			video.PreviewPath = settings.NO_IMAGE_FOUND_FILE
		}
		video.Tags = currentVideoTags
		video.ShowTitle = currentVideoShow
		videos = append(videos, video)
	}

	results, err = dbConnection.Query(settings.GET_ALL_VIDEOS_COUNT_QUERY)
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		}, response
	}

	for results.Next() {
		results.Scan(
			&numberOfVideos,
		)
	}

	response.IsSuccessful = true

	return models.VideosResponse{
		Videos:         videos,
		NumberOfVideos: numberOfVideos,
		VideoTags:      allVideoTags,
	}, response
}

func GetShows() ([]models.Show, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	shows := []models.Show{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	results, err := dbConnection.Query(settings.GET_SHOWS_QUERY)
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return shows, response
	}

	for results.Next() {
		var show = models.Show{}

		err = results.Scan(
			&show.ID,
			&show.Title,
			&show.PreviewPath,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}

		shows = append(shows, show)
	}
	response.IsSuccessful = true

	return shows, response
}

func GetVideoTags() ([]models.VideoTag, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videoTags := []models.VideoTag{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	results, err := dbConnection.Query(settings.GET_VIDEO_TAGS_QUERY)
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return videoTags, response
	}

	for results.Next() {

		var videoTag = models.VideoTag{}

		err := results.Scan(
			&videoTag.ID,
			&videoTag.Title,
			&videoTag.VideoID,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}

		videoTags = append(videoTags, videoTag)
	}
	response.IsSuccessful = true

	return videoTags, response
}

func GetVideoTagsByVideoID(videoID uuid.UUID) ([]models.VideoTag, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videoTags := []models.VideoTag{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	results, err := dbConnection.Query(settings.GET_VIDEO_TAGS_BY_VIDEO_ID_QUERY + "'" + videoID.String() + "'")
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return videoTags, response
	}

	for results.Next() {

		var videoTag = models.VideoTag{}

		err := results.Scan(
			&videoTag.ID,
			&videoTag.Title,
			&videoTag.VideoID,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}

		videoTags = append(videoTags, videoTag)
	}
	response.IsSuccessful = true

	return videoTags, response
}

func GetVideoByID(videoID uuid.UUID) (models.Video, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	video := models.Video{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}

	results, err := dbConnection.Query(settings.GET_VIDEO_BY_ID_QUERY + "'" + videoID.String() + "'")
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return video, response
	}

	for results.Next() {
		err := results.Scan(
			&video.ID,
			&video.Title,
			&video.Type,
			&video.Episode,
			&video.ShowID,
			&video.Duration,
			&video.Language,
			&video.PreviewPath,
			&video.VideoPath,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}
	}

	if !utilities.IsValidImageFile(video.PreviewPath) {
		video.PreviewPath = settings.NO_IMAGE_FOUND_FILE
	}

	response.IsSuccessful = true
	return video, response
}

func GetVideosByType(
	videoType string,
	limitValue string,
	offsetValue string,
) (models.VideosResponse, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videos := []models.Video{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	numberOfVideos := 0
	allVideoTags, videoTagsResponse := GetVideoTags()
	if !videoTagsResponse.IsSuccessful {
		response.Message = "Error getting tags " + videoTagsResponse.Message
		return models.VideosResponse{
				Videos:    videos,
				VideoTags: allVideoTags,
			},
			response
	}

	query := settings.GET_VIDEOS_BY_TYPE_QUERY + "'" + videoType + "'"

	if limitValue != "" && offsetValue != "" {
		query += " LIMIT " + limitValue + " OFFSET " + offsetValue
	} else if limitValue != "" {
		query += query + " LIMIT " + limitValue
	} else if offsetValue != "" {
		query += query + " OFFSET " + offsetValue
	}

	results, err := dbConnection.Query(query)
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return models.VideosResponse{
				Videos:    videos,
				VideoTags: allVideoTags,
			},
			response
	}

	for results.Next() {
		var video = models.Video{}

		err := results.Scan(
			&video.ID,
			&video.Title,
			&video.Type,
			&video.Episode,
			&video.ShowID,
			&video.Duration,
			&video.Language,
			&video.PreviewPath,
			&video.VideoPath,
		)
		if err != nil {
			response.Message = "Error scanning row: " + err.Error()
		}
		currentVideoTags := []string{}
		for i := 0; i < len(allVideoTags); i++ {
			if video.ID == allVideoTags[i].VideoID {
				currentVideoTags = append(currentVideoTags, allVideoTags[i].Title)
			}
		}
		if !utilities.IsValidImageFile(video.PreviewPath) {
			video.PreviewPath = settings.NO_IMAGE_FOUND_FILE
		}
		video.Tags = currentVideoTags
		videos = append(videos, video)
	}

	results, err = dbConnection.Query(settings.GET_VIDEOS_BY_TYPE_QUERY + "'" + videoType + "'")
	if err != nil {
		response.Message = "Error running query: " + err.Error()
		return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		}, response
	}

	for results.Next() {
		results.Scan(
			&numberOfVideos,
		)
	}

	response.IsSuccessful = true

	return models.VideosResponse{
			Videos:         videos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideoTags,
		},
		response
}

func GetVideosBySearch(searchTerm string, limitValue string, offsetValue string) (models.VideosResponse, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videosWithSearchTerm := []models.Video{}
	returnedVideos := []models.Video{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	numberOfVideos := 0
	allVideos, allVideosResponse := GetAllVideos("", "")
	if !allVideosResponse.IsSuccessful {
		response.Message = "Error getting all videos: " + allVideosResponse.Message
		return models.VideosResponse{
			Videos:         returnedVideos,
			NumberOfVideos: numberOfVideos,
			VideoTags:      allVideos.VideoTags,
		}, response
	}

	for i := 0; i < len(allVideos.Videos); i++ {
		currentVideo := allVideos.Videos[i]
		if utilities.IsSearchTermInVideo(searchTerm, currentVideo) {
			videosWithSearchTerm = append(videosWithSearchTerm, currentVideo)
		}
	}

	if limitValue != "" && offsetValue != "" {
		limitValueInt, err := strconv.Atoi(limitValue)
		if err != nil {
			response.Message = "Limit is not an int: " + allVideosResponse.Message
			return models.VideosResponse{
				Videos:         returnedVideos,
				NumberOfVideos: numberOfVideos,
				VideoTags:      allVideos.VideoTags,
			}, response
		}
		offsetValueInt, err := strconv.Atoi(offsetValue)
		if err != nil {
			response.Message = "Offset is not an int: " + allVideosResponse.Message
			return models.VideosResponse{
				Videos:         returnedVideos,
				NumberOfVideos: numberOfVideos,
				VideoTags:      allVideos.VideoTags,
			}, response
		}
		if limitValueInt > len(videosWithSearchTerm)-offsetValueInt {
			limitValueInt = len(videosWithSearchTerm) - offsetValueInt
		}
		for i := offsetValueInt; i < limitValueInt; i++ {
			if utilities.IsValidImageFile(videosWithSearchTerm[i].PreviewPath) {
				videosWithSearchTerm[i].PreviewPath = settings.NO_IMAGE_FOUND_FILE
			}
			returnedVideos = append(returnedVideos, videosWithSearchTerm[i])
			numberOfVideos++
		}
	} else if limitValue != "" {
		limitValueInt, err := strconv.Atoi(limitValue)
		if err != nil {
			response.Message = "Limit is not an int: " + allVideosResponse.Message
			return models.VideosResponse{
				Videos:         returnedVideos,
				NumberOfVideos: numberOfVideos,
				VideoTags:      allVideos.VideoTags,
			}, response
		}
		if limitValueInt > len(videosWithSearchTerm) {
			limitValueInt = len(videosWithSearchTerm)
		}
		for i := 0; i < limitValueInt; i++ {
			if utilities.IsValidImageFile(videosWithSearchTerm[i].PreviewPath) {
				videosWithSearchTerm[i].PreviewPath = settings.NO_IMAGE_FOUND_FILE
			}
			returnedVideos = append(returnedVideos, videosWithSearchTerm[i])
			numberOfVideos++
		}

	} else if offsetValue != "" {
		offsetValueInt, err := strconv.Atoi(limitValue)
		if err != nil {
			response.Message = "Offset is not an int: " + allVideosResponse.Message
			return models.VideosResponse{
				Videos:         returnedVideos,
				NumberOfVideos: numberOfVideos,
				VideoTags:      allVideos.VideoTags,
			}, response
		}
		for i := offsetValueInt; i < len(videosWithSearchTerm); i++ {
			if utilities.IsValidImageFile(videosWithSearchTerm[i].PreviewPath) {
				videosWithSearchTerm[i].PreviewPath = settings.NO_IMAGE_FOUND_FILE
			}
			returnedVideos = append(returnedVideos, videosWithSearchTerm[i])
			numberOfVideos++
		}
	} else {
		returnedVideos = videosWithSearchTerm
	}

	response.IsSuccessful = true

	return models.VideosResponse{
		Videos:         returnedVideos,
		NumberOfVideos: numberOfVideos,
		VideoTags:      allVideos.VideoTags,
	}, response
}

func GetVideosByShow(showID string) ([]models.Video, models.ServiceResponse) {
	dbConnection := services.ConnectToDB()
	defer dbConnection.Close()
	videos := []models.Video{}
	response := models.ServiceResponse{
		IsSuccessful: false,
		Message:      "",
	}
	allVideoTags, videoTagsResponse := GetVideoTags()
	if !videoTagsResponse.IsSuccessful {
		response.Message = "Error getting tags " + videoTagsResponse.Message
		return videos, response
	}
	allShows, showsResponse := GetVideoTags()
	if !showsResponse.IsSuccessful {
		response.Message = "Error getting tags " + showsResponse.Message
		return videos, response
	}

	results, err := dbConnection.Query(settings.GET_VIDEOS_BY_SHOW_ID_QUERY + "'" + showID + "'")
	if err != nil {
		response.Message = "Error running query " + err.Error()
		return videos, response
	}

	for results.Next() {
		var video = models.Video{}

		err := results.Scan(
			&video.ID,
			&video.Title,
			&video.Type,
			&video.Episode,
			&video.ShowID,
			&video.Duration,
			&video.Language,
			&video.PreviewPath,
			&video.VideoPath,
		)
		if err != nil {
			response.Message = "Error scanning row " + err.Error()
		}
		currentVideoTags := []string{}
		for i := 0; i < len(allVideoTags); i++ {
			if video.ID == allVideoTags[i].VideoID {
				currentVideoTags = append(currentVideoTags, allVideoTags[i].Title)
			}
		}
		currentVideoShow := ""
		for i := 0; i < len(allShows); i++ {
			if *video.ShowID == allShows[i].ID {
				currentVideoShow = allShows[i].Title
			}
		}
		if !utilities.IsValidImageFile(video.PreviewPath) {
			video.PreviewPath = settings.NO_IMAGE_FOUND_FILE
		}
		video.Tags = currentVideoTags
		video.ShowTitle = currentVideoShow
		videos = append(videos, video)
	}
	response.IsSuccessful = true

	return videos, response
}
