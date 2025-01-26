export default interface IVideoOld {
    ID: string;
    Title: string;
    Type: string;
    Episode: string;
    ShowID: number | null;
    Duration: string;
    Language: string;
    Tags: string[];
    PreviewPath: string;
    VideoPath: string
}