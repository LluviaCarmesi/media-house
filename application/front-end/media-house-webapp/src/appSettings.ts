// URIs
export const API_SERVER_URI = "http://192.168.1.247:445/api/";
export const VIDEOS_SERVER_URI = "http://192.168.1.247:444";
export const VIDEOS_URI = API_SERVER_URI + "videos/";
export const VIDEOS_SEARCH_URI = VIDEOS_URI + "search/";
export const MOVIES_URI = VIDEOS_URI + "movies/";
export const SHOWS_URI = VIDEOS_URI + "shows/";

// values
export const CHUNK_SIZE = 50 * 1024 * 1024; // 50MB
export const DEFAULT_LIMIT = 50;
export const DEFAULT_OFFSET = 50;
export const DEFAULT_NUMBER_OF_PAGES_TO_SHOW = 10;

// choices
export const TYPES = {
    SHOW: {
        value: "show",
        label: "Show"
    },
    MOVIE: {
        value: "movie",
        label: "Movie"
    }
};
export const TYPES_OPTIONS = [
    TYPES.MOVIE,
    TYPES.SHOW
]