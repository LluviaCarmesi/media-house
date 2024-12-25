package models

type ServiceResponse struct {
	IsSuccessful bool   `json:is_successful`
	ErrorMessage string `json:error_message`
}

type Video struct {
	// database attributes
	ID          string `json:id`
	Title       string `json:title`
	Type        string `json:type`
	Episode     string `json:episode`
	ShowID      int    `json:show_id`
	Duration    string `json:duration`
	Language    string `json:language`
	PreviewPath string `json:preview_path`
	VideoPath   string `json:video_path`
	// custom attributes
	PreviewFile []byte   `json:preview_path`
	VideoFile   []byte   `json:video_file`
	Tags        []string `json:tags`
	ShowTitle   string   `json:show_title`
}

type VideoChunks struct {
	VideoChunkNumber int `json:video_chunk_number`
	videoTotalChunks int `json:video_total_chunks`
}

type VideoTag struct {
	ID      int    `json:id`
	Title   string `json:title`
	VideoID string `json:video_id`
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
	Videos    []Video    `json:videos`
	VideoTags []VideoTag `json:video_tags`
}
