namespace media_house_api.Models
{
    public class VideosWithTags
    {
        private List<Video> videos = new List<Video>();
        private List<string> tags = new List<string>();


        public VideosWithTags(List<Video> videos, List<string> tags)
        {
            this.videos = videos;
            this.tags = tags;
        }

        public List<Video> Videos
        {
            get
            {
                return videos;
            }
            set
            {
                videos = value;
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
    }
}
