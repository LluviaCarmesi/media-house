<script lang="ts">
    import "../styles/pages/home.css";
    import "../styles/common.css";
    import Navigation from "../components/Navigation.svelte";
    import Filters from "../components/Filters.svelte";
    import type IVideo from "../interfaces/IVideo";
    import getVideosBySearch from "../services/getVideosBySearch";
    import {
        DEFAULT_LIMIT,
        DEFAULT_OFFSET,
        VIDEOS_SERVER_URI,
    } from "../appSettings";
    import Pager from "../components/Pager.svelte";
    import getVideos from "../services/getVideos";
    import { onMount } from "svelte";
    import { afterNavigate } from "$app/navigation";

    let videos: IVideo[] = [];
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

    function checkPageState() {
        const URLParameters = new URLSearchParams(window.location.search);
        const searchTermValue = URLParameters.get("q");
        if (!!searchTermValue) {
            searchTerm = searchTermValue;
        }
        const pageNumberValue = URLParameters.get("page_number");
        if (!!pageNumberValue) {
            currentPage = parseInt(pageNumberValue);
        } else {
            currentPage = 1;
        }

        if (!!searchTerm) {
            getVideosBySearchResponse();
        } else {
            getVideosResponse();
        }
    }

    onMount(() => {
        checkPageState();
    });

    afterNavigate(() => {
        checkPageState();
    });

    async function getVideosResponse() {
        isLoading = true;
        errorMessage = "";
        const videoResponse = await getVideos(
            DEFAULT_LIMIT,
            DEFAULT_OFFSET * (currentPage - 1),
        );
        if (videoResponse.isSuccessful) {
            videos = videoResponse.videos;
            numberOfVideos = videoResponse.numberOfVideos;
            maxPages = Math.ceil(numberOfVideos / DEFAULT_LIMIT);
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }

    async function getVideosBySearchResponse() {
        if (!searchTerm) {
            getVideosResponse();
            return;
        }
        isLoading = true;
        errorMessage = "";
        const videoResponse = await getVideosBySearch(
            searchTerm,
            DEFAULT_LIMIT,
            DEFAULT_OFFSET * (currentPage - 1),
        );
        if (videoResponse.isSuccessful) {
            videos = videoResponse.videos;
            numberOfVideos = videoResponse.numberOfVideos;
            maxPages = Math.ceil(numberOfVideos / DEFAULT_LIMIT);
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }

    function addSearchQueryParameter() {
        window.open(`?q=${searchTerm}`, "_self");
    }
</script>

<Navigation />

<div>
    <div class="textContainer">
        <h2 class="mainHeading">Media House</h2>
    </div>
    <Filters
        label={"Search Content"}
        onChangeTextbox={getSearchTerm}
        onEnterTextbox={addSearchQueryParameter}
        onChangeTagsDropdown={null}
        tagsDropdownOptions={[]}
        currentSearchValue={searchTerm}
    />
    <div class="videosContainer">
        {#if isLoading}
            <span>{loadingMessage}</span>
        {:else if !!errorMessage}
            <span>{errorMessage}</span>
        {:else}
            <div class="imageGrid">
                {#each videos as video}
                    <a href={"/video/" + video.ID}>
                        <img
                            src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`}
                        />
                        <p>{video.Title}</p>
                    </a>
                {/each}
            </div>
        {/if}
    </div>
</div>

{#if !isLoading}
    <Pager {currentPage} {maxPages} pageQueryString={"?page_number"} />
{/if}
