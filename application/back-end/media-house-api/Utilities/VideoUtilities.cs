using media_house_api.Models;
using Newtonsoft.Json;

namespace media_house_api.Utilities
{
    public static class VideoUtilities
    {
        public static List<Video> ConvertJSONToModel(String JSON)
        {
            return JsonConvert.DeserializeObject<List<Video>>(JSON);
        }
    }
}
