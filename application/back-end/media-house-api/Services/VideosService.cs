using media_house_api.Models;
using media_house_api.Utilities;
using Microsoft.AspNetCore.Mvc;

namespace media_house_api.Services
{
    public static class VideosService
    {
        public static IActionResult GetVideoJSON()
        {
            String JSON = String.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                String line = streamReader.ReadLine();
                while (line != null)
                {
                    JSON+= line;
                    line = streamReader.ReadLine();
                }
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new {error = "Couldn't get JSON file"});
            }

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

        public static IActionResult GetVideoById(String id)
        {
            String JSON = String.Empty;
            StreamReader streamReader = new StreamReader(AppSettings.videosJSONPath);
            try
            {
                String line = streamReader.ReadLine();
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

            if (foundVideo == null) {
                return new BadRequestObjectResult(new { error = "No Video Found"});
            }

            return new OkObjectResult(foundVideo);
        }
    }
}
