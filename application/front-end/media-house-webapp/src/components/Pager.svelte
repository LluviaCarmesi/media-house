<script lang="ts">
    import { DEFAULT_NUMBER_OF_PAGES_TO_SHOW } from "../appSettings";
    import "../styles/components/pager.css";

    export let maxPages;
    export let currentPage;
    export let pageQueryString;

    const pagesToShow: number[] = [currentPage];

    for (let i = 1; i < 2*DEFAULT_NUMBER_OF_PAGES_TO_SHOW; i++) {
        if ((currentPage - i) > 0) {
            pagesToShow.push(currentPage - i)
        }
        if ((currentPage + i) <= maxPages) {
            pagesToShow.push(currentPage + i)
        }
        if (pagesToShow.length > 10) {
            break;
        }
    }

    pagesToShow.sort((a, b) => {return a - b});
</script>

<div class="pagerContainer">
    {#if currentPage - 1 > 0}
        <a href={`${pageQueryString}=${currentPage - 1}`}>&lt</a>
    {/if}
    {#each pagesToShow as pageToShow}
        {#if pageToShow === currentPage}
            <a class="currentPage" href={`${pageQueryString}=${pageToShow}`}>{pageToShow}</a>
        {:else}
            <a href={`${pageQueryString}=${pageToShow}`}>{pageToShow}</a>
        {/if}
    {/each}
    {#if currentPage + 1 <= maxPages}
        <a href={`${pageQueryString}=${currentPage + 1}`}>&gt</a>
    {/if}
</div>