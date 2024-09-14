using media_house_api.Models;
using media_house_api.Utilities;
using Microsoft.AspNetCore.Mvc;
using MySqlConnector;

namespace media_house_api.Services
{
    public static class ShowsServices
    {
        private static MySqlConnection mySqlConnection = new MySqlConnection(AppSettings.SQL_CONNECTION_STRING);

        public static IActionResult GetShows()
        {
            List<Show> shows = new List<Show>();
            MySqlCommand mySqlCommand;
            try
            {
                mySqlConnection.Open();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.SHOWS_TABLE}", mySqlConnection);
                MySqlDataReader showsReader = mySqlCommand.ExecuteReader();
                while (showsReader.Read())
                {
                    shows.Add
                        (
                        new Show
                            (
                            showsReader.GetInt32(0),
                            showsReader.GetString(1)
                            )
                        );
                }
                showsReader.Close();
            }
            catch (MySqlException exception)
            {
                return new BadRequestObjectResult(new { error = exception.Message });
            }
            catch (Exception exception)
            {
                return new BadRequestObjectResult(new { error = exception.Message });
            }
            finally
            {
                mySqlConnection.Close();
            }

            if (shows.Count == 0)
            {
                return new BadRequestObjectResult(new { error = "No Shows Found" });
            }

            return new OkObjectResult(shows);
        }
    }
}
