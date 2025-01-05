import isStatusGood from "../utilities/isStatusGood";
import type INewVideo from "../interfaces/IVideo";
import { VIDEOS_URI } from "../appSettings";

export default async function addVideo(
    item: INewVideo,
    videoFileChunkNumber: number,
    videoFileTotalChunks: number,
    chunk: Blob
) {
    let doesErrorExist = false;
    let errorMessage = "";
    const bodyData = new FormData();
    bodyData.append("PreviewFile", item.PreviewFile);
    bodyData.append("Title", item.Title);
    bodyData.append("Type", item.Type);
    bodyData.append("Episode", item.Episode);
    bodyData.append("ShowID", item.ShowID);
    bodyData.append("Duration", item.Duration);
    bodyData.append("Tags", JSON.stringify(item.Tags));
    bodyData.append("Language", item.Language);
    bodyData.append("VideoFileChunkNumber", videoFileChunkNumber.toString());
    bodyData.append("VideoFileTotalChunks", videoFileTotalChunks.toString());
    bodyData.append("VideoFile", chunk, item.VideoFile.name);

    await fetch(VIDEOS_URI, {
        method: "POST",
        body: bodyData
    })
        .then((response) => {
            doesErrorExist = !isStatusGood(response.status);
            return response.json();
        })
        .then((result) => {
            if (doesErrorExist) {
                errorMessage = result.response;
            }
        })
        .catch((error) => {
            doesErrorExist = true;
            errorMessage = error;
            console.log(error);
        });

    return { doesErrorExist, errorMessage };
}