import * as SETTINGS from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";
import * as Strings from "../strings/ENUSStrings";
import type IShow from "../interfaces/IShow";

export default async function getShows() {
    const returnedResponse: {
        shows: IShow[],
        isSuccessful: boolean,
        errorMessage: string
    }
        = {
        shows: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${SETTINGS.SHOWS_API_URI}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.shows = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.shows.length === 0) {
        returnedResponse.errorMessage = Strings.NoVideosErrorMessage;
    }
    return returnedResponse;
}