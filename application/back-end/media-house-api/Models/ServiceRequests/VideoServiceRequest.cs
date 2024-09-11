using Invoicer.Models;
using media_house_api.Models;

namespace media_house_api.Models.ServiceRequests
{
    public class VideoServiceRequest : ServiceRequest
    {
        public Video video = new Video();

        public VideoServiceRequest(bool isSuccessful, string result, Video video) : base(isSuccessful, result)
        {
            this.video = video;
        }

        public Video Video
        {
            get
            {
                return video;
            }
            set
            {
                video = value;
            }
        }
    }
}
