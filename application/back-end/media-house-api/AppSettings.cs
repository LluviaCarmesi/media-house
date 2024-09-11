namespace media_house_api
{
    internal static class AppSettings
    {
        internal static string SQL_CONNECTION_STRING = "server=192.168.1.248;database=private_frio_content;uid=private_frio_content_user;pwd=media_house;";
        //internal static string VIDEOS_DIRECTORY = "C:\\Users\\meradil\\Videos";
        //internal static string PHOTOS_DIRECTORY = "C:\\Users\\meradil\\Pictures";
        internal static string VIDEOS_DIRECTORY = "/var/www/private-frio.com/public/videos";
        internal static string PHOTOS_DIRECTORY = "/var/www/private-frio.com/public/photos";
        internal static string LOCAL_HOST_TESTING_DOMAIN = "http://localhost:5173";
        internal static string LOCAL_HOST_PRODUCTION_DOMAIN = "http://192.168.1.248:690";
        internal static string SHOWS_TABLE = "shows";
        internal static string VIDEOS_TABLE = "videos";
        internal static string VIDEO_TAGS_TABLE = "video_tags";
        internal static string PRIMARY_ID_COLUMN = "id";
        internal static string VIDEO_ID_COLUMN = "video_id";
        internal static string ADD_VIDEO_COLUMNS =
            "id, title, type, episode, show_id, duration, language, preview_path, video_path";
        internal static string ADD_VIDEO_TAGS_COLUMNS = "title, video_id";
        internal static string TYPE_COLUMN = "type";
    }
}
