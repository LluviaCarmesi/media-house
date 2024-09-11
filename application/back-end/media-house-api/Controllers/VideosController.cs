using Invoicer.Models.ServiceRequests;
using media_house_api.Models.ServiceRequests;
using media_house_api.Services;
using media_house_api.Utilities;
using Microsoft.AspNetCore.Mvc;

namespace media_house_api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class VideosController : ControllerBase
    {

        [HttpGet("")]
        public IActionResult GetVideos()
        {

            return VideosServices.GetVideos();
        }

        [HttpGet("Search/{searchTerm}")]
        public IActionResult GetVideosBySearch(string searchTerm)
        {
            return VideosServices.GetVideoBySearch(searchTerm);
        }

        [HttpGet("{id}")]
        public IActionResult GetVideoById(string id)
        {
            return VideosServices.GetVideoById(id);
        }

        [HttpGet("Type/{type}")]
        public IActionResult GetVideoByType(string type)
        {
            return VideosServices.GetVideoByType(type);
        }

        [HttpPost("Add")]
        public IActionResult AddVideo
        (
            [FromForm] IFormFile? previewFile,
            [FromForm] IFormFile? videoFile,
            [FromForm] string? title,
            [FromForm] string? type,
            [FromForm] string? episode,
            [FromForm] string? showID,
            [FromForm] string? duration,
            [FromForm] string? language,
            [FromForm] string? tags
        )
        {
            VideoServiceRequest videoModelValidationResult =
                VideosUtilities.CheckVideoModel
                (
                    previewFile,
                    videoFile,
                    title,
                    type,
                    episode,
                    showID,
                    duration,
                    language,
                    tags
                );
            if (!videoModelValidationResult.IsSuccessful)
            {
                return BadRequest(new { error = videoModelValidationResult.Result });
            }
            Task<CommonServiceRequest> videoUploadValidation = VideosServices.AddVideoFiles(previewFile, videoFile);
            CommonServiceRequest videoUploadValidationResponse = videoUploadValidation.Result;
            if (!videoUploadValidationResponse.IsSuccessful)
            {
                return BadRequest(new { error = videoUploadValidationResponse.Result });
            }
            videoModelValidationResult.Video.VideoPath = $"{AppSettings.VIDEOS_DIRECTORY}/{videoFile.FileName}";
            videoModelValidationResult.Video.PreviewPath = $"{AppSettings.PHOTOS_DIRECTORY}/{previewFile.FileName}";

            CommonServiceRequest videoAddValidation = VideosServices.AddVideo(videoModelValidationResult.Video);
            if (!videoAddValidation.IsSuccessful)
            {
                return BadRequest(new { error = videoAddValidation.Result });
            }
            return new OkObjectResult(new { response = videoAddValidation.Result });
        }
    }
}
