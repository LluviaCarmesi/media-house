export default interface IVideo {
    id: string;
    title: string;
    type: string;
    episode: string;
    showID: number | null;
    duration: string;
    language: string;
    tags: string[];
    previewPath: string;
    gifPreviewPath: string;
    videoPath: string
}