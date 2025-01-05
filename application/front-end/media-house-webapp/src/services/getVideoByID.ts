import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IVideo from "../interfaces/IVideoOld";
import { VIDEOS_URI } from "../appSettings";

export default async function getVideoByID(id: string) {
    const returnedResponse: {
        video: IVideo,
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        video: {
            ID: "",
            Title: "",
            Type: "",
            Episode: "",
            ShowID: 0,
            Language: "",
            Duration: "",
            Tags: [""],
            PreviewPath: "",
            VideoPath: "",
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${VIDEOS_URI}${id}`)
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
    if (!returnedResponse.video.ID) {
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    return returnedResponse;
}