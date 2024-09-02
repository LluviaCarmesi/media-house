using media_house_api.Models;
using Newtonsoft.Json;

namespace media_house_api.Utilities
{
    public static class VideoUtilities
    {
        public static List<Video> ConvertJSONToModel(string JSON)
        {
            return JsonConvert.DeserializeObject<List<Video>>(JSON);
        }

        public static List<string> GetAllTagsFromVideos(List<Video> videos)
        {
            List<string> gatheredTags = new List<string>();
            for (int i = 0; i < videos.Count; i++)
            {
                Video currentVideo = videos[i];
                if (currentVideo.Tags.Count == 0)
                {
                    continue;
                }
                for (int j = 0; j < currentVideo.Tags.Count; j++)
                {
                    string currentTag = currentVideo.Tags[j];
                    if (string.IsNullOrEmpty(gatheredTags.FirstOrDefault(
                        tag => tag.ToLower() == currentTag.ToLower()
                        )))
                    {
                        gatheredTags.Add(currentTag);
                    }
                }
            }
            return gatheredTags;
        }
    }
}
