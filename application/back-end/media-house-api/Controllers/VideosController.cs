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

        [HttpGet("Search")]
        public IActionResult GetVideosBySearch()
        {
            return new OkObjectResult("");
        }

        [HttpGet("{id}")]
        public IActionResult GetVideoById(String id)
        {
            return VideosService.GetVideoById(id);
        }
    }
}
