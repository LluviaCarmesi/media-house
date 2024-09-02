import * as SETTINGS from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IVideo from "../interfaces/IVideo";

export default async function getVideosByType(type: string) {
    const returnedResponse: {
        videos: IVideo[],
        tags: string[],
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        videos: [],
        tags: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${SETTINGS.VIDEOS_TYPE_API_URI}/${type}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.videos = result.videos;
                returnedResponse.tags = result.tags;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.videos.length === 0) {
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    return returnedResponse;
}