namespace media_house_api.Models
{
    public class VideoTags
    {
        private int id = 0;
        private string title = string.Empty;
        private Guid videoID = Guid.Empty;

        public VideoTags(int id, string title, Guid videoID)
        {
            this.id = id;
            this.title = title;
            this.videoID = videoID;
        }

        public int Id
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

        public Guid VideoID
        {
            get
            {
                return videoID;
            }
            set
            {
                videoID = value;
                videoID = value;
            }
        }
    }
}