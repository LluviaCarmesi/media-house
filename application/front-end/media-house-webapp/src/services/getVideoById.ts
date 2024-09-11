import * as SETTINGS from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IVideo from "../interfaces/IVideo";

export default async function getVideoById(id: string) {
    const returnedResponse: {
        video: IVideo,
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        video: {
            id: "",
            title: "",
            type: "",
            episode: "",
            showID: 0,
            duration: "",
            tags: [""],
            previewPath: "",
            gifPreviewPath: "",
            videoPath: "",
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${SETTINGS.VIDEOS_API_URI}/${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.video = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.video.id) {
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    return returnedResponse;
}