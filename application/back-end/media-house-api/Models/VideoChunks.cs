namespace media_house_api.Models
{
    public class VideoChunks
    {
        private int videoFileChunkNumber = 0;
        private int videoFileTotalChunks = 0;

        public VideoChunks()
        {

        }

        public VideoChunks(int videoFileChunkNumber, int videoFileTotalChunks)
        {
            this.videoFileChunkNumber = videoFileChunkNumber;
            this.videoFileTotalChunks = videoFileTotalChunks;
        }

        public int VideoFileChunkNumber
        {
            get
            {
                return videoFileChunkNumber;
            }
            set
            {
                videoFileChunkNumber = value;
            }
        }

        public int VideoFileTotalChunks
        {
            get
            {
                return videoFileTotalChunks;
            }
            set
            {
                videoFileTotalChunks = value;
            }
        }
    }
}
