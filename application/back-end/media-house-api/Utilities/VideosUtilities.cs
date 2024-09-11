using Invoicer.Models.ServiceRequests;
using media_house_api.Models;
using media_house_api.Models.ServiceRequests;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using System.Text;


namespace media_house_api.Utilities
{
    public static class VideosUtilities
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
        public static List<Video> JoinTagsWithVideos(List<Video> videos, List<VideoTags> videoTags)
        {
            for (int i = 0; i < videos.Count; i++)
            {
                List<string> tagsTitles = new List<string>();
                List<VideoTags> foundTags =
                    videoTags.FindAll(tag => tag.VideoID == videos[i].Id);
                for (int j = 0; j < foundTags.Count; j++)
                {
                    tagsTitles.Add(foundTags[j].Title);
                }
                videos[i].Tags = tagsTitles;
            }
            return videos;
        }
        public static VideoServiceRequest CheckVideoModel
        (
            IFormFile previewFile,
            IFormFile videoFile,
            string title,
            string type,
            string episode,
            string showID,
            string duration,
            string language,
            string tags
        )
        {
            Video video = new Video();
            bool isValid = true;
            string result = string.Empty;
            if (previewFile == null || previewFile.Length == 0)
            {
                isValid = false;
                result = "previewFile is required";
            }
            if (videoFile == null || videoFile.Length == 0)
            {
                isValid = false;
                result = "videoFile is required";
            }

            video.Id = Guid.NewGuid();

            // title validation
            if (string.IsNullOrEmpty(title))
            {
                isValid = false;
                result = "Title is required";
            }
            else
            {
                video.Title = title;
            }

            // type validation
            if (string.IsNullOrEmpty(type))
            {
                isValid = false;
                result = "Type is required";
            }
            else
            {
                video.Type = type;
            }

            // episode validation
            if (type == "movie")
            {
                video.Episode = "";
            }
            else if (string.IsNullOrEmpty(episode))
            {
                isValid = false;
                result = "Episode is required";
            }
            else
            {
                video.Episode = episode;
            }

            // showID validation
            int showIDInt = 0;
            if (type == "movie")
            {
                video.ShowID = null;
            }
            else if (string.IsNullOrEmpty(showID))
            {
                isValid = false;
                result = "Show ID is required";
            }
            else if (int.TryParse(showID, out showIDInt))
            {
                video.ShowID = showIDInt;
            }

            // duration validation
            if (string.IsNullOrEmpty(duration))
            {
                isValid = false;
                result = "Duration is required";
            }
            else
            {
                video.Duration = duration;
            }

            // language validation
            if (string.IsNullOrEmpty(language))
            {
                isValid = false;
                result = "Language is required";
            }
            else
            {
                video.Language = language;
            }

            
            video.Tags = JsonConvert.DeserializeObject<List<string>>(tags);

            return new VideoServiceRequest(isValid, result, video);
        }
    }
}
