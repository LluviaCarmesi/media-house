<script lang="ts">
    import "../../styles/common.css";
    import Navigation from "../../components/Navigation.svelte";
    import TextField from "../../components/TextField.svelte";
    import Dropdown from "../../components/Dropdown.svelte";
    import FilePicker from "../../components/FilePicker.svelte";
    import { CHUNK_SIZE, TYPES_OPTIONS } from "../../appSettings";
    import getShows from "../../services/getShows";
    import type IDropdownOption from "../../interfaces/IDropdownOption";
    import addVideo from "../../services/addVideo";
    import type INewVideo from "../../interfaces/INewVideo";

    let showDropdownOptions: IDropdownOption[] = [];
    let isLoading = true;
    let errorMessage = "";
    let isSuccessful = false;

    const newVideoItem: INewVideo = {
        title: "",
        type: "",
        episode: "",
        showID: "",
        language: "",
        duration: "",
        tagsString: "",
        tags: [],
        previewFile: null,
        videoFile: null,
    };

    let progress = "";

    async function getShowsResponse() {
        const showsResponse = await getShows();
        if (showsResponse.isSuccessful) {
            for (let i = 0; i < showsResponse.shows.length; i++) {
                showDropdownOptions.push({
                    value: showsResponse.shows[i].id,
                    label: showsResponse.shows[i].title,
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
        newVideoItem.tags = newVideoItem.tagsString.split(",");

        const videoFileTotalChunks = Math.ceil(
            newVideoItem.videoFile.size / CHUNK_SIZE,
        );
        for (let i = 0; i < newVideoItem.videoFile.size; i += CHUNK_SIZE) {
            const end = Math.min(i + CHUNK_SIZE, newVideoItem.videoFile.size);
            const chunk = newVideoItem.videoFile.slice(i, end);
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
                inputID="title"
                currentValue={newVideoItem.title}
                fieldLabel="Title"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <Dropdown
                inputID="type"
                currentValue={newVideoItem.type}
                fieldLabel="Type"
                onDropdownChange={handleDropdownChange}
                dropdownOptions={TYPES_OPTIONS}
            />
        </div>
        {#if newVideoItem.type == "show"}
            <div class="centerContainer">
                <TextField
                    inputID="episode"
                    currentValue={newVideoItem.episode}
                    fieldLabel="Episode Number"
                    onChangeTextField={handleTextChange}
                />
            </div>
            <div class="centerContainer">
                <Dropdown
                    inputID="showID"
                    currentValue={newVideoItem.showID}
                    fieldLabel="Show"
                    onDropdownChange={handleDropdownChange}
                    dropdownOptions={[]}
                />
            </div>
        {/if}
        <div class="centerContainer">
            <TextField
                inputID="language"
                currentValue={newVideoItem.language}
                fieldLabel="Language"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="duration"
                currentValue={newVideoItem.duration}
                fieldLabel="Duration"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="tagsString"
                currentValue={newVideoItem.tagsString}
                fieldLabel="Tags"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="previewFile"
                fieldLabel="Attach Preview"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="videoFile"
                fieldLabel="Attach Video"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="endContainer">
            <button on:click={submitVideo}>Submit</button>
        </div>
    </div>
</div>
