using media_house_api.Models;
using media_house_api.Utilities;
using Microsoft.AspNetCore.Mvc;

namespace media_house_api.Services
{
    public static class VideosService
    {
        public static IActionResult GetVideoJSON()
        {
            string JSON = string.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                string line = streamReader.ReadLine();
                while (line != null)
                {
                    JSON += line;
                    line = streamReader.ReadLine();
                }
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't get JSON file" });
            }
            streamReader.Close();

            List<Video> videos = new List<Video>();
            try
            {
                videos = VideoUtilities.ConvertJSONToModel(JSON);
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't properly deserialize JSON file" });
            }

            return new OkObjectResult(videos);
        }

        public static IActionResult GetVideoBySearch(string searchTerm)
        {
            string JSON = string.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                string line = streamReader.ReadLine();
                while (line != null)
                {
                    JSON += line;
                    line = streamReader.ReadLine();
                }
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't get JSON file" });
            }
            streamReader.Close();

            List<Video> videos = new List<Video>();
            try
            {
                videos = VideoUtilities.ConvertJSONToModel(JSON);
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't properly deserialize JSON file" });
            }
            List<Video> foundVideos = new List<Video>();
            foundVideos = videos.FindAll(
                video => video.Title.ToLower().Contains(searchTerm.ToLower()) ||
                (
                    video.Tags.Find(tag => tag.ToLower() == searchTerm.ToLower()) != null
                )
            );
            if (foundVideos.Count == 0)
            {
                return new BadRequestObjectResult(new { error = "No Videos Found" });
            }
            return new OkObjectResult(foundVideos);
        }

        public static IActionResult GetVideoById(string id)
        {
            string JSON = string.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                string line = streamReader.ReadLine();
                while (line != null)
                {
                    JSON += line;
                    line = streamReader.ReadLine();
                }
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't get JSON file" });
            }
            streamReader.Close();

            List<Video> videos = new List<Video>();
            try
            {
                videos = VideoUtilities.ConvertJSONToModel(JSON);
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't properly deserialize JSON file" });
            }
            Video foundVideo = videos.FirstOrDefault(video => video.Id.ToLower() == id.ToLower());

            if (foundVideo == null)
            {
                return new BadRequestObjectResult(new { error = "No Videos Found" });
            }

            return new OkObjectResult(foundVideo);
        }

        public static IActionResult GetVideoByType(string type)
        {
            string JSON = string.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                string line = streamReader.ReadLine();
                while (line != null)
                {
                    JSON += line;
                    line = streamReader.ReadLine();
                }
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't get JSON file" });
            }
            streamReader.Close();

            List<Video> videos = new List<Video>();
            try
            {
                videos = VideoUtilities.ConvertJSONToModel(JSON);
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new { error = "Couldn't properly deserialize JSON file" });
            }
            List<Video> foundVideos = new List<Video>();
            foundVideos = videos.FindAll(video => video.Type.ToLower() == type.ToLower());

            if (foundVideos.Count == 0)
            {
                return new BadRequestObjectResult(new { error = "No Videos Found" });
            }

            VideosWithTags videosWithTags = new VideosWithTags(
                foundVideos,
                VideoUtilities.GetAllTagsFromVideos(foundVideos)
            );

            return new OkObjectResult(videosWithTags);
        }
    }
}
