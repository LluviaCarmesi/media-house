using media_house_api.Services;
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

            return VideosService.GetVideoJSON();
        }

        [HttpGet("Search/{searchTerm}")]
        public IActionResult GetVideosBySearch(string searchTerm)
        {
            return VideosService.GetVideoBySearch(searchTerm);
        }

        [HttpGet("{id}")]
        public IActionResult GetVideoById(string id)
        {
            return VideosService.GetVideoById(id);
        }

        [HttpGet("Type/{type}")]
        public IActionResult GetVideoByType(string type)
        {
            return VideosService.GetVideoByType(type);
        }
    }
}
