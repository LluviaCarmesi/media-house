<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import Filters from "../../components/Filters.svelte";
    import type IVideo from "../../interfaces/IVideo";
    import getVideosByType from "../../services/getVideosByType";
    import "../../styles/pages/movies.css";
    import "../../styles/common.css";
    import getFilteredVideos from "../../utilities/getFilteredVideos";
    import separateVideos from "../../utilities/separateVideos";

    let videos: IVideo[] = [];
    let firstVideos: IVideo[] = [];
    let secondVideos: IVideo[] = [];
    let thirdVideos: IVideo[] = [];
    let tags: string[] = [];
    let filteredVideos: IVideo[] = [];
    let isLoading = true;
    const loadingMessage = "Loading...";
    let searchTerm = "";
    let chosenTag = "";

    let errorMessage = "";
    async function getVideosByTypeResponse() {
        const videoResponse = await getVideosByType("movie");
        if (videoResponse.isSuccessful) {
            videos = videoResponse.videos;
            tags = videoResponse.tags;
            filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
            filteredVideos = videos;
            const separatedVideos = separateVideos(filteredVideos, 3);
            firstVideos = separatedVideos[0];
            secondVideos = separatedVideos[1];
            thirdVideos = separatedVideos[2];
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }
    getVideosByTypeResponse();

    function getSearchTerm(event: any) {
        searchTerm = event.target.value.toLowerCase();
        filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
        const separatedVideos = separateVideos(filteredVideos, 3);
        firstVideos = separatedVideos[0];
        secondVideos = separatedVideos[1];
        thirdVideos = separatedVideos[2];
    }

    function getChosenTag(event: any) {
        const chosenTagValue = event.target.value.toLowerCase();
        if (chosenTagValue == "any") {
            chosenTag = "";
        } else {
            chosenTag = chosenTagValue;
        }
        filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
        const separatedVideos = separateVideos(filteredVideos, 3);
        firstVideos = separatedVideos[0];
        secondVideos = separatedVideos[1];
        thirdVideos = separatedVideos[2];
    }
</script>

<Navigation />
<Filters
    label={"Search Movies"}
    onChangeTextbox={getSearchTerm}
    onChangeTagsDropdown={getChosenTag}
    tagsDropdownOptions={tags}
/>

<div class="videosContainer">
    {#if isLoading}
        <span>{loadingMessage}</span>
    {:else if !!errorMessage}
        <span>{errorMessage}</span>
    {:else}
        <div class="column">
            {#each firstVideos as video}
                <a href={"/video/" + video.id}>
                    <img src={video.previewPath} />
                    <p>{video.title}</p>
                </a>
            {/each}
        </div>
        <div class="column">
            {#each secondVideos as video}
                <a href={"/video/" + video.id}>
                    <img src={video.previewPath} />
                    <p>{video.title}</p>
                </a>
            {/each}
        </div>
        <div class="column">
            {#each thirdVideos as video}
                <a href={"/video/" + video.id}>
                    <img src={video.previewPath} />
                    <p>{video.title}</p>
                </a>
            {/each}
        </div>
    {/if}
</div>
