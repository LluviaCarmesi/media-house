namespace media_house_api.Models
{
    public class Video
    {
        private string id = string.Empty;
        private string title = string.Empty;
        private string type = string.Empty;
        private List<string> tags = new List<string>();
        private string previewPath = string.Empty;
        private string videoPath = string.Empty;

        public Video(
            string id,
            string title,
            List<string> tags,
            string previewPath,
            string videoPath
            )
        {
            this.id = id;
            this.title = title;
            this.tags = tags;
            this.previewPath = previewPath;
            this.videoPath = videoPath;
        }

        public string Id
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
