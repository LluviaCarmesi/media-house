<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import Filters from "../../components/Filters.svelte";
    import type IVideo from "../../interfaces/IVideo";
    import getVideosByType from "../../services/getVideosByType";
    import "../../styles/pages/movies.css";
    import "../../styles/common.css";
    import getFilteredVideos from "../../utilities/getFilteredVideos";
    import separateVideos from "../../utilities/separateVideos";
    import { DEFAULT_LIMIT, DEFAULT_OFFSET, VIDEOS_SERVER_URI } from "../../appSettings";
    import createDropdownOptions from "../../utilities/createDropdownOptions";
    import type IDropdownOption from "../../interfaces/IDropdownOption";
    import Pager from "../../components/Pager.svelte";
    import { onMount } from "svelte";

    let videos: IVideo[] = [];
    let firstVideos: IVideo[] = [];
    let secondVideos: IVideo[] = [];
    let thirdVideos: IVideo[] = [];
    let fourthVideos: IVideo[] = [];
    let fifthVideos: IVideo[] = [];
    let tagsDropdownOptions: IDropdownOption[] = [];
    let filteredVideos: IVideo[] = [];
    let isLoading = true;
    const loadingMessage = "Loading...";
    let searchTerm = "";
    let chosenTag = "";
    let currentPage = 1;
    let maxPages = 1;
    let numberOfVideos = 0;

    onMount(() => {
        const URLParameters = new URLSearchParams(window.location.search);
        const pageNumberValue = URLParameters.get("page_number");
        if (!!pageNumberValue) {
            currentPage = parseInt(pageNumberValue);
        }
    });

    let errorMessage = "";
    async function getVideosByTypeResponse() {
        const videoResponse = await getVideosByType("movies", DEFAULT_LIMIT, DEFAULT_OFFSET * (currentPage - 1));
        if (videoResponse.isSuccessful) {
            videos = videoResponse.videos;
            numberOfVideos = videoResponse.numberOfVideos;
            maxPages = Math.ceil(numberOfVideos / DEFAULT_LIMIT); 
            tagsDropdownOptions = createDropdownOptions(videoResponse.tags);
            filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
            filteredVideos = videos;
            const separatedVideos = separateVideos(filteredVideos, 5);
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
    getVideosByTypeResponse();

    function getSearchTerm(event: any) {
        searchTerm = event.target.value.toLowerCase();
        filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
        const separatedVideos = separateVideos(filteredVideos, 5);
        firstVideos = separatedVideos[0];
        secondVideos = separatedVideos[1];
        thirdVideos = separatedVideos[2];
        fourthVideos = separatedVideos[3];
        fifthVideos = separatedVideos[4];
    }

    function getChosenTag(event: any) {
        const chosenTagValue = event.target.value.toLowerCase();
        if (chosenTagValue == "any") {
            chosenTag = "";
        } else {
            chosenTag = chosenTagValue;
        }
        filteredVideos = getFilteredVideos(videos, searchTerm, chosenTag);
        const separatedVideos = separateVideos(filteredVideos, 5);
        firstVideos = separatedVideos[0];
        secondVideos = separatedVideos[1];
        thirdVideos = separatedVideos[2];
        fourthVideos = separatedVideos[3];
        fifthVideos = separatedVideos[4];
    }
</script>

<Navigation />
<Filters
    label={"Search Movies"}
    onChangeTextbox={getSearchTerm}
    onChangeTagsDropdown={getChosenTag}
    tagsDropdownOptions={tagsDropdownOptions}
/>

<div class="videosContainer">
    {#if isLoading}
        <span>{loadingMessage}</span>
    {:else if !!errorMessage}
        <span>{errorMessage}</span>
    {:else}
    <table>
        <tbody>
            <tr>
                <td></td>
            </tr>
        </tbody>
    </table>
        <div class="column">
            {#each firstVideos as video}
                <a href={"/video/" + video.ID}>
                    <img
                        alt=""
                        src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`}
                    />
                    <p>{video.Title}</p>
                </a>
            {/each}
        </div>
        <div class="column">
            {#each secondVideos as video}
                <a href={"/video/" + video.ID}>
                    <img
                        alt=""
                        src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`}
                    />
                    <p>{video.Title}</p>
                </a>
            {/each}
        </div>
        <div class="column">
            {#each thirdVideos as video}
                <a href={"/video/" + video.ID}>
                    <img
                        alt=""
                        src={`${VIDEOS_SERVER_URI}/${video.PreviewPath}`}
                    />
                    <p>{video.Title}</p>
                </a>
            {/each}
        </div>
    {/if}
</div>

<Pager
    currentPage={currentPage}
    maxPages={maxPages}
    pageQueryString={"?page_number"}
/>