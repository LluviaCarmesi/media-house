import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IVideo from "../interfaces/IVideoOld";
import { VIDEOS_URI } from "../appSettings";

export default async function getVideosByType(type: string, limitValue: number, offsetValue: number) {
    const returnedResponse: {
        videos: IVideo[],
        numberOfVideos: number,
        tags: string[],
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        videos: [],
        numberOfVideos: 0,
        tags: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${VIDEOS_URI}${type}/?limit=${limitValue}&offset=${offsetValue}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.videos = result.Videos;
                returnedResponse.numberOfVideos = result.NumberOfVideos;
                returnedResponse.tags = result.VideoTags;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.videos.length === 0) {
        returnedResponse.isSuccessful = false;
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    return returnedResponse;
}