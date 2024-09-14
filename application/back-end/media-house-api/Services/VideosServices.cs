using Invoicer.Models.ServiceRequests;
using media_house_api.Models;
using media_house_api.Utilities;
using Microsoft.AspNetCore.Mvc;
using MySqlConnector;
using System;

namespace media_house_api.Services
{
    public static class VideosServices
    {
        private static MySqlConnection mySqlConnection = new MySqlConnection(AppSettings.SQL_CONNECTION_STRING);

        public static IActionResult GetVideos()
        {
            List<Video> videos = new List<Video>();
            List<VideoTags> videoTags = new List<VideoTags>();
            MySqlCommand mySqlCommand;
            try
            {
                mySqlConnection.Open();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.VIDEOS_TABLE}", mySqlConnection);
                MySqlDataReader videosReader = mySqlCommand.ExecuteReader();
                while (videosReader.Read())
                {
                    videos.Add
                        (
                        new Video
                            (
                            videosReader.GetGuid(0),
                            videosReader.GetString(1),
                            videosReader.GetString(2),
                            videosReader.GetString(3),
                            videosReader["show_id"] as int?,
                            videosReader.GetString(5),
                            videosReader.GetString(6),
                            new List<string>(),
                            videosReader.GetString(7),
                            videosReader.GetString(8)
                            )
                        );
                }
                videosReader.Close();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.VIDEO_TAGS_TABLE}", mySqlConnection);
                MySqlDataReader videoTagsReader = mySqlCommand.ExecuteReader();
                while (videoTagsReader.Read())
                {
                    videoTags.Add
                        (
                        new VideoTags
                            (
                            videoTagsReader.GetInt32(0),
                            videoTagsReader.GetString(1),
                            videoTagsReader.GetGuid(2)
                            )
                        );
                }
                videoTagsReader.Close();
                videos = VideosUtilities.JoinTagsWithVideos(videos, videoTags);
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

            if (videos.Count == 0)
            {
                return new BadRequestObjectResult(new { error = "No Videos Found" });
            }
            return new OkObjectResult(videos);
        }

        public static IActionResult GetVideoBySearch(string searchTerm)
        {
            List<Video> videos = new List<Video>();
            List<VideoTags> videoTags = new List<VideoTags>();
            MySqlCommand mySqlCommand;
            try
            {
                mySqlConnection.Open();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.VIDEOS_TABLE}", mySqlConnection);
                MySqlDataReader videosReader = mySqlCommand.ExecuteReader();
                while (videosReader.Read())
                {
                    videos.Add
                        (
                        new Video
                            (
                            videosReader.GetGuid(0),
                            videosReader.GetString(1),
                            videosReader.GetString(2),
                            videosReader.GetString(3),
                            videosReader["show_id"] as int?,
                            videosReader.GetString(5),
                            videosReader.GetString(6),
                            new List<string>(),
                            videosReader.GetString(7),
                            videosReader.GetString(8)
                            )
                        );
                }
                videosReader.Close();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.VIDEO_TAGS_TABLE}", mySqlConnection);
                MySqlDataReader videoTagsReader = mySqlCommand.ExecuteReader();
                while (videoTagsReader.Read())
                {
                    videoTags.Add
                        (
                        new VideoTags
                            (
                            videoTagsReader.GetInt32(0),
                            videoTagsReader.GetString(1),
                            videoTagsReader.GetGuid(2)
                            )
                        );
                }
                videoTagsReader.Close();
                videos = VideosUtilities.JoinTagsWithVideos(videos, videoTags);
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
            Video foundVideo = new Video();
            List<string> foundTags = new List<string>();
            MySqlCommand mySqlCommand;
            try
            {
                mySqlConnection.Open();
                mySqlCommand = new MySqlCommand(
                    $"SELECT * FROM {AppSettings.VIDEOS_TABLE} " +
                    $"WHERE {AppSettings.PRIMARY_ID_COLUMN} = '{id}'",
                    mySqlConnection
                    );
                MySqlDataReader videosReader = mySqlCommand.ExecuteReader();
                while (videosReader.Read())
                {
                    foundVideo.Id = videosReader.GetGuid(0);
                    foundVideo.Title = videosReader.GetString(1);
                    foundVideo.Type = videosReader.GetString(2);
                    foundVideo.Episode = videosReader.GetString(3);
                    foundVideo.ShowID = videosReader["show_id"] as int?;
                    foundVideo.Duration = videosReader.GetString(5);
                    foundVideo.Language = videosReader.GetString(6);
                    new List<string>();
                    foundVideo.PreviewPath = videosReader.GetString(7);
                    foundVideo.VideoPath = videosReader.GetString(8);
                }
                videosReader.Close();

                if (foundVideo.Id == Guid.Empty)
                {
                    return new BadRequestObjectResult(new { error = "No Videos Found" });
                }

                mySqlCommand = new MySqlCommand(
                    $"SELECT * FROM {AppSettings.VIDEO_TAGS_TABLE} " +
                    $"WHERE {AppSettings.VIDEO_ID_COLUMN} = '{id}'",
                    mySqlConnection
                );
                MySqlDataReader videoTagsReader = mySqlCommand.ExecuteReader();
                while (videoTagsReader.Read())
                {
                    foundTags.Add
                        (
                            videoTagsReader.GetString(1)
                        );
                }
                foundVideo.Tags = foundTags;
                videoTagsReader.Close();
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
            return new OkObjectResult(foundVideo);
        }

        public static IActionResult GetVideoByType(string type)
        {
            List<Video> videos = new List<Video>();
            List<VideoTags> videoTags = new List<VideoTags>();
            MySqlCommand mySqlCommand;
            try
            {
                mySqlConnection.Open();
                mySqlCommand = new MySqlCommand(
                    $"SELECT * FROM {AppSettings.VIDEOS_TABLE} " +
                    $"WHERE {AppSettings.TYPE_COLUMN} = '{type}'",
                    mySqlConnection
                );
                MySqlDataReader videosReader = mySqlCommand.ExecuteReader();
                while (videosReader.Read())
                {
                    videos.Add
                        (
                        new Video
                            (
                            videosReader.GetGuid(0),
                            videosReader.GetString(1),
                            videosReader.GetString(2),
                            videosReader.GetString(3),
                            videosReader["show_id"] as int?,
                            videosReader.GetString(5),
                            videosReader.GetString(6),
                            new List<string>(),
                            videosReader.GetString(7),
                            videosReader.GetString(8)
                            )
                        );
                }
                videosReader.Close();
                mySqlCommand = new MySqlCommand($"SELECT * FROM {AppSettings.VIDEO_TAGS_TABLE}", mySqlConnection);
                MySqlDataReader videoTagsReader = mySqlCommand.ExecuteReader();
                while (videoTagsReader.Read())
                {
                    videoTags.Add
                        (
                        new VideoTags
                            (
                            videoTagsReader.GetInt32(0),
                            videoTagsReader.GetString(1),
                            videoTagsReader.GetGuid(2)
                            )
                        );
                }
                videoTagsReader.Close();
                videos = VideosUtilities.JoinTagsWithVideos(videos, videoTags);
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

            if (videos.Count == 0)
            {
                return new BadRequestObjectResult(new { error = "No Videos Found" });
            }

            VideosWithTags videosWithTags = new VideosWithTags(
                videos,
                VideosUtilities.GetAllTagsFromVideos(videos)
            );

            return new OkObjectResult(videosWithTags);
        }

        public static CommonServiceRequest AddVideo(Video video)
        {

            bool isSuccessful = true;
            string result = string.Empty;
            mySqlConnection.Open();
            MySqlCommand mySqlAddVideoCommand;
            mySqlAddVideoCommand = new MySqlCommand
                (
                $"INSERT INTO {AppSettings.VIDEOS_TABLE} " +
                $"({AppSettings.ADD_VIDEO_COLUMNS}) " +
                $"VALUES (@id, @title, @type, @episode, @showID, @duration, @language, @previewPath, @videoPath)",
                mySqlConnection
                );
            try
            {
                mySqlAddVideoCommand.Parameters.Add("@id", MySqlDbType.VarChar).Value = video.Id;
                mySqlAddVideoCommand.Parameters.Add("@title", MySqlDbType.VarChar).Value = video.Title;
                mySqlAddVideoCommand.Parameters.Add("@type", MySqlDbType.VarChar).Value = video.Type;
                mySqlAddVideoCommand.Parameters.Add("@episode", MySqlDbType.VarChar).Value = video.Episode;
                mySqlAddVideoCommand.Parameters.Add("@showID", MySqlDbType.Int32).Value = video.ShowID;
                mySqlAddVideoCommand.Parameters.Add("@duration", MySqlDbType.VarChar).Value = video.Duration;
                mySqlAddVideoCommand.Parameters.Add("@language", MySqlDbType.VarChar).Value = video.Language;
                mySqlAddVideoCommand.Parameters.Add("@previewPath", MySqlDbType.VarChar).Value = video.PreviewPath;
                mySqlAddVideoCommand.Parameters.Add("@videoPath", MySqlDbType.VarChar).Value = video.VideoPath;
                mySqlAddVideoCommand.ExecuteNonQuery();

                for (int i = 0; i < video.Tags.Count; i++)
                {
                    MySqlCommand mySqlAddVideoTagsCommand = new MySqlCommand
                    (
                    $"INSERT INTO {AppSettings.VIDEO_TAGS_TABLE} " +
                    $"({AppSettings.ADD_VIDEO_TAGS_COLUMNS}) " +
                    $"VALUES (@title, @video_id)",
                    mySqlConnection
                    );

                    mySqlAddVideoTagsCommand.Parameters.Add("@title", MySqlDbType.VarChar).Value = video.Tags[i];
                    mySqlAddVideoTagsCommand.Parameters.Add("@video_id", MySqlDbType.VarChar).Value = video.Id;
                    mySqlAddVideoTagsCommand.ExecuteNonQuery();
                }
                isSuccessful = true;
                result = "Video attributes were added";
            }
            catch (MySqlException exception)
            {
                result = "Video attributes couldn't be added" + exception.Message;
                isSuccessful = false;
            }
            catch (Exception exception)
            {
                result = "Video attributes couldn't be added" + exception.Message;
                isSuccessful = false;
            }
            finally
            {
                mySqlConnection.Close();
            }

            return new CommonServiceRequest(isSuccessful, result);
        }

        public static async Task<CommonServiceRequest> AddVideoFiles(
            IFormFile previewFile,
            IFormFile videoFile,
            int videoFileChunkNumber,
            int videoFileTotalChunks
        )
        {
            Console.WriteLine(videoFileChunkNumber);
            if (
                File.Exists($"{AppSettings.VIDEOS_DIRECTORY}/{videoFile.FileName}") ||
                File.Exists($"{AppSettings.PHOTOS_DIRECTORY}/{previewFile.FileName}")
                )
            {
                return new CommonServiceRequest(false, "The file(s) already exist");
            }
            FileStream videoFileChunkStream = new FileStream
            (
                $"{AppSettings.VIDEOS_TEMP_DIRECTORY}/{videoFile.FileName}_{videoFileChunkNumber}",
                FileMode.Create
            );
            try
            {
                await videoFile.CopyToAsync(videoFileChunkStream);
                videoFileChunkStream.Close();
            }
            catch (Exception exception)
            {
                return new CommonServiceRequest(false, exception.Message);
            }

            if (videoFileChunkNumber != videoFileTotalChunks - 1)
            {
                return new CommonServiceRequest(true, "Chunks were uploaded");
            }

            FileStream videoFileStream = new FileStream
            (
                $"{AppSettings.VIDEOS_DIRECTORY}/{videoFile.FileName}",
                FileMode.Create
            );

            FileStream previewFileStream = new FileStream
            (
                $"{AppSettings.PHOTOS_DIRECTORY}/{previewFile.FileName}",
                FileMode.Create
            );

            try
            {
                await previewFile.CopyToAsync(previewFileStream);
                previewFileStream.Close();

                for (int i = 0; i < videoFileTotalChunks; i++)
                {
                    string tempPath =
                        $"{AppSettings.VIDEOS_TEMP_DIRECTORY}/{videoFile.FileName}_{i}";

                    FileStream tempStream = new FileStream(tempPath, FileMode.Open);
                    await tempStream.CopyToAsync(videoFileStream);
                    tempStream.Close();
                    File.Delete(tempPath);
                }
                videoFileStream.Close();

            }
            catch (Exception exception)
            {
                return new CommonServiceRequest(false, exception.Message);
            }
            return new CommonServiceRequest(true, "Files successfully uploaded");
        }
    }
}
