import type IVideo from "../interfaces/IVideo";

export default function separateVideos(videos: IVideo[], amount: number) {
    let separatedVideos: IVideo[][] = [];

    for (let i = 0; i < amount; i++) {
        separatedVideos[i] = [];
    }
    for (let i = 0; i < videos.length; i++) {
        separatedVideos[i % amount].push(videos[i]);
    }

    return separatedVideos;
}