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
    import type INewVideo from "../../interfaces/INewVideo";

    let showDropdownOptions: IDropdownOption[] = [];
    let isLoading = true;
    let errorMessage = "";
    let isSuccessful = false;

    const newVideoItem: any = {
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
        newVideoItem[event.target.id as keyof INewVideo] =
            event.target.files[0];
    }

    function handleTextChange(event: any) {
        newVideoItem[event.target.id as keyof INewVideo] = event.target.value;
    }

    function handleDropdownChange(event: any) {
        newVideoItem[event.target.id as keyof INewVideo] = event.target.value;
    }

    async function submitVideo() {
        newVideoItem.Tags = newVideoItem.TagsString.split(",");

        const videoFileTotalChunks = Math.ceil(
            newVideoItem.VideoFile.size / CHUNK_SIZE,
        );
        for (let i = 0; i < newVideoItem.VideoFile.size; i += CHUNK_SIZE) {
            const end = Math.min(i + CHUNK_SIZE, newVideoItem.VideoFile.size);
            const chunk = newVideoItem.VideoFile.slice(i, end);
            const videoFileChunkNumber = Math.floor(i / CHUNK_SIZE);
            progress = (
                (videoFileChunkNumber / videoFileTotalChunks) *
                100
            ).toFixed(2);
            const addVideoResponse = await addVideo(
                newVideoItem,
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
                currentValue={newVideoItem.Title}
                fieldLabel="Title"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <Dropdown
                inputID="Type"
                currentValue={newVideoItem.Type}
                fieldLabel="Type"
                onDropdownChange={handleDropdownChange}
                dropdownOptions={TYPES_OPTIONS}
            />
        </div>
        {#if newVideoItem.Type == TYPES.SHOW.value}
            <div class="centerContainer">
                <TextField
                    inputID="Episode"
                    currentValue={newVideoItem.Episode}
                    fieldLabel="Episode Number"
                    onChangeTextField={handleTextChange}
                />
            </div>
            <div class="centerContainer">
                <Dropdown
                    inputID="ShowID"
                    currentValue={newVideoItem.ShowID}
                    fieldLabel="Show"
                    onDropdownChange={handleDropdownChange}
                    dropdownOptions={[]}
                />
            </div>
        {/if}
        <div class="centerContainer">
            <TextField
                inputID="Language"
                currentValue={newVideoItem.Language}
                fieldLabel="Language"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="Duration"
                currentValue={newVideoItem.Duration}
                fieldLabel="Duration"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="TagsString"
                currentValue={newVideoItem.TagsString}
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
