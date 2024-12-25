import * as SETTINGS from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IVideo from "../interfaces/IVideo";

export default async function getVideosBySearch(searchTerm: string) {
    const returnedResponse: {
        videos: IVideo[],
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        videos: [],
        isSuccessful: false,
        errorMessage: ""
    }

    /*
    await fetch(`${SETTINGS.VIDEOS_SEARCH_API_URI}/${searchTerm}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.videos = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.videos.length === 0) {
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    */
    return returnedResponse;
}