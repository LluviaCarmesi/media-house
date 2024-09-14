using Invoicer.Models;
using media_house_api.Models;

namespace media_house_api.Models.ServiceRequests
{
    public class VideoChunksServiceRequest : ServiceRequest
    {
        public VideoChunks videoChunks = new VideoChunks();

        public VideoChunksServiceRequest(bool isSuccessful, string result, VideoChunks videoChunks) : base(isSuccessful, result)
        {
            this.videoChunks = videoChunks;
        }

        public VideoChunks VideoChunks
        {
            get
            {
                return videoChunks;
            }
            set
            {
                videoChunks = value;
            }
        }
    }
}
