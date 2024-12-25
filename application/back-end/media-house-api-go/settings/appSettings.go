package settings

// paths
const API_ROOT = "/api/"
const VIDEOS_PATH = API_ROOT + "videos/"
const MOVIES_PATH = VIDEOS_PATH + "movies/"
const SHOWS_PATH = VIDEOS_PATH + "shows/"

// queries
const GET_VIDEO_BY_ID_QUERY = "select * from videos WHERE id = "
const GET_VIDEOS_BY_TYPE_QUERY = "select * from videos WHERE type = "
const GET_VIDEOS_BY_SHOW_ID_QUERY = "select * from videos WHERE show_id = "
const GET_VIDEO_TAGS_QUERY = "select * from video_tags;"
const GET_SHOWS_QUERY = "select * from shows;"

// types
const SHOWS_TYPE = "show"
const MOVIES_TYPE = "movie"
