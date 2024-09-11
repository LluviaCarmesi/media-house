using media_house_api.Services;
using Microsoft.AspNetCore.Mvc;

namespace media_house_api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ShowsController: ControllerBase
    {
        [HttpGet("")]
        public IActionResult GetShows()
        {

            return ShowsServices.GetShows();
        }
    }
}
