<script lang="ts">
    import "../styles/pages/home.css";
    import "../styles/common.css";
    import Navigation from "../components/Navigation.svelte";
    import Filters from "../components/Filters.svelte";
    import type IVideo from "../interfaces/IVideo";
    import getVideosBySearch from "../services/getVideosBySearch";
    import separateVideos from "../utilities/separateVideos";
    import { VIDEOS_SERVER_URI } from "../appSettings";
    import Pager from "../components/Pager.svelte";

    let videos: IVideo[] = [];
    let firstVideos: IVideo[] = [];
    let secondVideos: IVideo[] = [];
    let thirdVideos: IVideo[] = [];
    let fourthVideos: IVideo[] = [];
    let fifthVideos: IVideo[] = [];
    let searchTerm = "";
    let errorMessage = "";
    let isLoading = false;
    const loadingMessage = "Loading...";
    let currentPage = 1;
    let maxPages = 1;
    let numberOfVideos = 0;

    function getSearchTerm(event: any) {
        searchTerm = event.target.value.toLowerCase();
    }

    async function getVideosBySearchResponse() {
        isLoading = true;
        errorMessage = "";
        const videoResponse = await getVideosBySearch(searchTerm);
        if (videoResponse.isSuccessful) {
            videos = videoResponse.videos;
            const separatedVideos = separateVideos(videos, 5);
            firstVideos = separatedVideos[0];
            secondVideos = separatedVideos[1];
            thirdVideos = separatedVideos[2];
            fourthVideos = separatedVideos[3];
            fifthVideos = separatedVideos[4];
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }
</script>

<Navigation />

<div>
    <div class="textContainer">
        <h2 class="mainHeading">Media House</h2>
    </div>
    <div class="textContainer">
        <span class="description">
            Use the links on top to watch your favorite content!
        </span>
    </div>
    <Filters
        label={"Search Content"}
        onChangeTextbox={getSearchTerm}
        onEnterTextbox={getVideosBySearchResponse}
        onChangeTagsDropdown={null}
        tagsDropdownOptions={[]}
    />
    <div class="videosContainer">
        {#if isLoading}
            <span>{loadingMessage}</span>
        {:else if !!errorMessage}
            <span>{errorMessage}</span>
        {:else}
            <div class="column">
                {#each firstVideos as video}
                    <a href={"/video/" + video.ID}>
                        <img src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`} />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
            <div class="column">
                {#each secondVideos as video}
                    <a href={"/video/" + video.ID}>
                        <img src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`} />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
            <div class="column">
                {#each thirdVideos as video}
                    <a href={"/video/" + video.ID}>
                        <img src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`} />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
            <div class="column">
                {#each fourthVideos as video}
                    <a href={"/video/" + video.ID}>
                        <img src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`} />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
            <div class="column">
                {#each fifthVideos as video}
                    <a href={"/video/" + video.ID}>
                        <img src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`} />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
        {/if}
    </div>
</div>

<Pager
    currentPage={currentPage}
    maxPages={maxPages}
    pageQueryString={"?page_number"}
/>