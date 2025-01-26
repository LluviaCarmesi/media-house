package settings

// ports
const SERVE_PORT = ":8080"

// URI paths
const API_ROOT = "/api/"
const VIDEOS_PATH = API_ROOT + "videos/"
const VIDEOS_SEARCH_PATH = VIDEOS_PATH + "search/"
const MOVIES_PATH = VIDEOS_PATH + "movies/"
const SHOWS_PATH = VIDEOS_PATH + "shows/"

// file directories
//const VIDEO_FILES_DIRECTORY = "/home/meradil/Downloads/"
//const VIDEO_FILES_SERVE_PATH = "Downloads/"
//const PREVIEW_FILES_DIRECTORY = "/home/meradil/Downloads/"
//const PREVIEW_FILES_SERVE_PATH = "Downloads/"
//const VIDEO_FILES_TEMP_DIRECTORY = "/home/meradil/Downloads/Temp/"

const VIDEO_FILES_DIRECTORY = "/media/vol_1/media_house/videos/"
const VIDEO_FILES_SERVE_PATH = "videos/"
const PREVIEW_FILES_DIRECTORY = "/media/vol_1/media_house/photos/"
const PREVIEW_FILES_SERVE_PATH = "photos/"
const VIDEO_FILES_TEMP_DIRECTORY = "/media/vol_1/temp/"

// queries
const GET_ALL_VIDEOS_QUERY = "SELECT * FROM videos"
const GET_ALL_VIDEOS_COUNT_QUERY = "SELECT COUNT(*) FROM videos;"
const GET_VIDEO_BY_ID_QUERY = "SELECT * FROM videos WHERE id = "
const GET_VIDEOS_BY_TYPE_QUERY = "SELECT * FROM videos WHERE type = "
const GET_VIDEOS_BY_TYPE_COUNT_QUERY = "SELECT COUNT(*) FROM videos WHERE type = "
const GET_VIDEOS_BY_SHOW_ID_QUERY = "SELECT * FROM videos WHERE show_id = "
const GET_VIDEO_TAGS_QUERY = "SELECT * FROM video_tags;"
const GET_VIDEO_TAGS_BY_VIDEO_ID_QUERY = "SELECT * FROM video_tags WHERE video_id = "
const GET_SHOWS_QUERY = "SELECT * FROM shows;"

const INSERT_SHOW_QUERY = "INSERT INTO shows (title,preview_path) VALUES(?,?)"
const INSERT_VIDEO_TAG_QUERY = "INSERT INTO video_tags (title,video_id) VALUES (?,?)"
const INSERT_VIDEO_QUERY = "INSERT INTO videos (id,title,type,episode,show_id,duration,language,preview_path,video_path) " +
	"VALUES (?,?,?,?,?,?,?,?,?)"

const UPDATE_SHOW_QUERY = "UPDATE shows SET title = ?, preview_path = ? WHERE id = ?"
const UPDATE_VIDEO_TAG_QUERY = "UPDATE shows SET title = ?, video_id = ? WHERE id = ?"
const UPDATE_VIDEO_QUERY = "UPDATE videos SET title = ?, type = ?, episode = ?, show_id = ?, " +
	"duration = ?, language = ?, preview_path = ?, video_path = ? WHERE id = ?"

const DELETE_SHOW_QUERY = "DELETE FROM shows WHERE id = ?"
const DELETE_VIDEO_TAG_QUERY = "DELETE FROM video_tags WHERE id = ?"
const DELETE_VIDEO_QUERY = "DELETE FROM videos WHERE id = ?"

// types
const SHOWS_TYPE = "show"
const MOVIES_TYPE = "movie"
