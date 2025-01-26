<script>
    import "../../../styles/pages/video.css";
    import Navigation from "../../../components/Navigation.svelte";
    export let data;
    import getVideoById from "../../../services/getVideoByID";
    import { VIDEOS_SERVER_URI } from "../../../appSettings";

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

    async function getVideoByIdResponse() {
        const videoResponse = await getVideoById(data.id);
        if (videoResponse.isSuccessful) {
            video = videoResponse.video;
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }
    getVideoByIdResponse();
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
    {/if}
</div>
