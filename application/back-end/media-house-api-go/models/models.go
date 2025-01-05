package models

import "github.com/google/uuid"

type ServiceResponse struct {
	IsSuccessful bool   `json:is_successful`
	Message      string `json:error_message`
}

type Video struct {
	// database attributes
	ID          uuid.UUID `json:id`
	Title       string    `json:title`
	Type        string    `json:type`
	Episode     string    `json:episode`
	ShowID      *int      `json:show_id`
	Duration    string    `json:duration`
	Language    string    `json:language`
	PreviewPath string    `json:preview_path`
	VideoPath   string    `json:video_path`
	// custom attributes
	PreviewFile     []byte   `json:preview_file`
	VideoFile       []byte   `json:video_file`
	PreviewFileName string   `json:preview_file_name`
	VideoFileName   string   `json:video_file_name`
	Tags            []string `json:tags`
	ShowTitle       string   `json:show_title`
}

type VideoChunks struct {
	VideoFileChunkNumber int `json:video_file_chunk_number`
	VideoFileTotalChunks int `json:video_file_total_chunks`
}

type VideoTag struct {
	ID      int       `json:id`
	Title   string    `json:title`
	VideoID uuid.UUID `json:video_id`
}

type Show struct {
	ID          int    `json:id`
	Title       string `json:title`
	PreviewPath string `json:preview_path`
}

type ModelCheckResponse struct {
	IsValid  bool   `json:is_valid`
	Response string `json:response`
}

type VideosResponse struct {
	Videos         []Video    `json:videos`
	NumberOfVideos int        `json:number_of_videos`
	VideoTags      []VideoTag `json:video_tags`
}
