import { VIDEOS_URI } from "../appSettings";
import isStatusGood from "../utilities/isStatusGood";

export default async function deleteVideoByID(id: string) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${VIDEOS_URI}${id}`, {
        method: "DELETE"
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