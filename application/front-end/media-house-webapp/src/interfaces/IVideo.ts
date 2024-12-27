export default interface IVideo {
    Id: string;
    Title: string;
    Type: string;
    Episode: string;
    ShowID: number | null;
    Duration: string;
    Language: string;
    Tags: string[];
    PreviewPath: string;
    GifPreviewPath: string;
    VideoPath: string
}