import * as SETTINGS from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";
import type INewVideo from "../interfaces/INewVideo";

export default async function addVideo(
    item: INewVideo,
    videoFileChunkNumber: number,
    videoFileTotalChunks: number,
    chunk: Blob
) {
    let doesErrorExist = false;
    let errorMessage = "";
    const bodyData = new FormData();
    bodyData.append("previewFile", item.previewFile);
    bodyData.append("title", item.title);
    bodyData.append("type", item.type);
    bodyData.append("episode", item.episode);
    bodyData.append("showID", item.showID);
    bodyData.append("duration", item.duration);
    bodyData.append("tags", JSON.stringify(item.tags));
    bodyData.append("language", item.language);
    bodyData.append("videoFileChunkNumber", videoFileChunkNumber.toString());
    bodyData.append("videoFileTotalChunks", videoFileTotalChunks.toString());
    bodyData.append("videoFile", chunk, item.videoFile.name);

    await fetch(`${SETTINGS.ADD_VIDEO_API_URI}`, {
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