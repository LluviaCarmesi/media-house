<script lang="ts">
    import "../../styles/common.css";
    import Navigation from "../../components/Navigation.svelte";
    import TextField from "../../components/TextField.svelte";
    import Dropdown from "../../components/Dropdown.svelte";
    import FilePicker from "../../components/FilePicker.svelte";
    import { CHUNK_SIZE, TYPES, TYPES_OPTIONS } from "../../appSettings";
    import getShows from "../../services/getShows";
    import type IDropdownOption from "../../interfaces/IDropdownOption";
    import addVideo from "../../services/addVideo";
    import type IVideo from "../../interfaces/IVideo";

    let showDropdownOptions: IDropdownOption[] = [];
    let isLoading = true;
    let errorMessage = "";
    let isSuccessful = false;

    const videoItem: any = {
        Title: "",
        Type: "",
        Episode: "",
        ShowID: 0,
        ShowTitle: "",
        Language: "",
        Duration: "",
        TagsString: "",
        Tags: [],
        VideoPath: "",
        PreviewPath: "",
        PreviewFile: null,
        VideoFile: null,
    };

    let progress = "";

    async function getShowsResponse() {
        const showsResponse = await getShows();
        if (showsResponse.isSuccessful) {
            for (let i = 0; i < showsResponse.shows.length; i++) {
                showDropdownOptions.push({
                    Value: showsResponse.shows[i].Id,
                    Label: showsResponse.shows[i].Title,
                });
            }
        } else {
            errorMessage = showsResponse.errorMessage;
        }
        isLoading = false;
    }

    getShowsResponse();

    function handleFileUpload(event: any) {
        videoItem[event.target.id as keyof IVideo] =
            event.target.files[0];
    }

    function handleTextChange(event: any) {
        videoItem[event.target.id as keyof IVideo] = event.target.value;
    }

    function handleDropdownChange(event: any) {
        videoItem[event.target.id as keyof IVideo] = event.target.value;
    }

    async function submitVideo() {
        videoItem.Tags = videoItem.TagsString.split(",");

        const videoFileTotalChunks = Math.ceil(
            videoItem.VideoFile.size / CHUNK_SIZE,
        );
        for (let i = 0; i < videoItem.VideoFile.size; i += CHUNK_SIZE) {
            const end = Math.min(i + CHUNK_SIZE, videoItem.VideoFile.size);
            const chunk = videoItem.VideoFile.slice(i, end);
            const videoFileChunkNumber = Math.floor(i / CHUNK_SIZE);
            progress = (
                (videoFileChunkNumber / videoFileTotalChunks) *
                100
            ).toFixed(2);
            const addVideoResponse = await addVideo(
                videoItem,
                videoFileChunkNumber,
                videoFileTotalChunks,
                chunk,
            );
            isSuccessful = !addVideoResponse.doesErrorExist;
            if (!isSuccessful) {
                errorMessage = "Couldn't add video";
                break;
            } else {
                errorMessage = "";
            }
        }
        progress = "";
    }
</script>

<Navigation />

<div>
    <div class="textContainer">
        <h2 class="mainHeading">Add a Video</h2>
    </div>
    <div class="textContainer">
        <span class="description">Add a video using the form below! </span>
    </div>
    {#if !!progress}
        <div class="textContainer">
            <span class="warning"
                >Adding the video! Don't leave! {progress}% complete</span
            >
        </div>
    {/if}
    {#if !!errorMessage}
        <div class="textContainer">
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="textContainer">
            <span class="success">Video was added</span>
        </div>
    {/if}
    <div class="form">
        <div class="centerContainer">
            <TextField
                inputID="Title"
                currentValue={videoItem.Title}
                fieldLabel="Title"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <Dropdown
                inputID="Type"
                currentValue={videoItem.Type}
                fieldLabel="Type"
                onDropdownChange={handleDropdownChange}
                dropdownOptions={TYPES_OPTIONS}
            />
        </div>
        {#if videoItem.Type == TYPES.SHOW.value}
            <div class="centerContainer">
                <TextField
                    inputID="Episode"
                    currentValue={videoItem.Episode}
                    fieldLabel="Episode Number"
                    onChangeTextField={handleTextChange}
                />
            </div>
            <div class="centerContainer">
                <Dropdown
                    inputID="ShowID"
                    currentValue={videoItem.ShowID}
                    fieldLabel="Show"
                    onDropdownChange={handleDropdownChange}
                    dropdownOptions={[]}
                />
            </div>
        {/if}
        <div class="centerContainer">
            <TextField
                inputID="Language"
                currentValue={videoItem.Language}
                fieldLabel="Language"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="Duration"
                currentValue={videoItem.Duration}
                fieldLabel="Duration"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="TagsString"
                currentValue={videoItem.TagsString}
                fieldLabel="Tags"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="PreviewFile"
                fieldLabel="Attach Preview"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="VideoFile"
                fieldLabel="Attach Video"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="endContainer">
            <button on:click={submitVideo}>Submit</button>
        </div>
    </div>
</div>
