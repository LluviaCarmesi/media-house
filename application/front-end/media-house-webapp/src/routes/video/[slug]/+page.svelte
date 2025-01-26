<script>
    import "../../../styles/pages/video.css";
    import "../../../styles/common.css";
    import Navigation from "../../../components/Navigation.svelte";
    export let data;
    import getVideoByID from "../../../services/getVideoByID";
    import { VIDEOS_SERVER_URI } from "../../../appSettings";
    import deleteVideoByID from "../../../services/deleteVideoByID";
    import { error } from "@sveltejs/kit";

    let video = {
        ID: "",
        Title: "",
        Type: "",
        Tags: [""],
        PreviewPath: "",
        VideoPath: "",
    };
    let isLoading = true;
    const loadingMessage = "Loading...";

    let errorMessage = "";

    async function getVideoByIDResponse() {
        const videoResponse = await getVideoByID(data.id);
        if (videoResponse.isSuccessful) {
            video = videoResponse.video;
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }

    async function deleteVideoByIDResponse() {
        errorMessage = "";
        const deleteVideoResponse = await deleteVideoByID(data.id);
        if (deleteVideoResponse.doesErrorExist) {
            errorMessage = deleteVideoResponse.errorMessage;
        }
        else {
            window.open("/", "_self");
        }
        isLoading = false;
    }
    getVideoByIDResponse();
</script>

<Navigation />
<div class="container">
    {#if isLoading}
        <h1>{loadingMessage}</h1>
    {:else}
        <h1>{video.Title}</h1>
        <h1>{errorMessage}</h1>
        <video controls>
            <source
                id="videoSource"
                src={`${VIDEOS_SERVER_URI}/${video.VideoPath}`}
                type="video/mp4"
            />
            Your browser doesn't support videos
            <track kind="captions" />
        </video>
        <div class="container">
            <button on:click={deleteVideoByIDResponse} class="deleteButton">Delete this video</button>
        </div>
    {/if}
</div>
