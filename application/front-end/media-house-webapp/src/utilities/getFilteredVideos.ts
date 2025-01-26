import type IVideo from "../interfaces/IVideoOld";

export default function (videos: IVideo[], searchTerm: string, chosenTag: string) {
    let filteredVideos = videos;
    if (!!searchTerm) {
        filteredVideos = filteredVideos.filter(
            (filteredVideo: { id: string; title: string, tags: string[] }) =>
                filteredVideo.title.toLowerCase().indexOf(searchTerm) !== -1,
        );
    }
    if (!!chosenTag) {
        filteredVideos = filteredVideos.filter(
            (filteredVideo: { id: string; title: string, tags: string[] }) =>
                filteredVideo.tags.filter((filteredVideoTag: string) =>
                    filteredVideoTag.toLowerCase() === chosenTag).length > 0
        );
    }

    /*
    if (tags.length > 0) {
        for (let i = 0; i < tags.length; i++) {
            filteredVideos = filteredVideos.filter(
                (filteredVideo: { id: string; title: string, tags: string[] }) =>
                    filteredVideo.tags.filter((tag: string) => tag === tags[i]).length > 0
            );
        }
    }
    */

    return filteredVideos;
}