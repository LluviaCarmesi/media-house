namespace media_house_api.Models
{
    public class Video
    {
        private Guid id = Guid.Empty;
        private string title = string.Empty;
        private string type = string.Empty;
        private string episode = string.Empty;
        private int? showID = 0;
        private string duration = string.Empty;
        private string language = string.Empty;
        private List<string> tags = new List<string>();
        private string previewPath = string.Empty;
        private string videoPath = string.Empty;

        public Video()
        {

        }
        public Video(
            Guid id,
            string title,
            string type,
            string episode,
            int? showID,
            string duration,
            string language,
            List<string> tags,
            string previewPath,
            string videoPath
            )
        {
            this.id = id;
            this.title = title;
            this.type = type;
            this.episode = episode;
            this.showID = showID;
            this.duration = duration;
            this.language = language;
            this.tags = tags;
            this.previewPath = previewPath;
            this.videoPath = videoPath;
        }

        public Guid Id
        {
            get
            {
                return id;
            }
            set
            {
                id = value;
            }
        }

        public string Title
        {
            get
            {
                return title;
            }
            set
            {
                title = value;
            }
        }

        public string Type
        {
            get
            {
                return type;
            }
            set
            {
                type = value;
            }
        }

        public string Episode
        {
            get
            {
                return episode;
            }
            set
            {
                episode = value;
            }
        }

        public int? ShowID
        {
            get
            {
                return showID;
            }
            set
            {
                showID = value;
            }
        }

        public string Duration
        {
            get
            {
                return duration;
            }
            set
            {
                duration = value;
            }
        }

        public string Language
        {
            get
            {
                return language;
            }
            set
            {
                language = value;
            }
        }

        public List<string> Tags
        {
            get
            {
                return tags;
            }
            set
            {
                tags = value;
            }
        }

        public string PreviewPath
        {
            get
            {
                return previewPath;
            }
            set
            {
                previewPath = value;
            }
        }

        public string VideoPath
        {
            get
            {
                return videoPath;
            }
            set
            {
                videoPath = value;
            }
        }
    }
}
