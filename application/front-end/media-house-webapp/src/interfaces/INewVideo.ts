export default interface INewVideo {
    title: string;
    type: string;
    episode: string;
    showID: string;
    duration: string;
    language: string;
    tagsString: string;
    tags: string[];
    previewFile: any;
    videoFile: any;
}