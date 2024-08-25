namespace media_house_api.Models
{
    public class Video
    {
        private string id = String.Empty;
        private String title = String.Empty;
        private List<String> tags = new List<String>();
        private String previewPath = String.Empty;
        private String videoPath = String.Empty;

        public Video(
            String id,
            String title,
            List<String> tags,
            String previewPath,
            String videoPath
            )
        {
            this.id = id;
            this.title = title;
            this.tags = tags;
            this.previewPath = previewPath;
            this.videoPath = videoPath;
        }

        public String Id
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

        public String Title
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

        public List<String> Tags
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

        public String PreviewPath
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

        public String VideoPath
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
