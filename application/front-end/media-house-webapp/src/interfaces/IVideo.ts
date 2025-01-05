export default interface IVideo {
    ID: string;
    Title: string;
    Type: string;
    Episode: string;
    ShowID: string;
    ShowTitle: string;
    Duration: string;
    Language: string;
    TagsString: string;
    Tags: string[];
    PreviewPath: string;
    VideoPath: string;
    PreviewFile: any;
    VideoFile: any;
}