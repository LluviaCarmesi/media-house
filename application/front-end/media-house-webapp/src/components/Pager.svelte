<script lang="ts">
    import { DEFAULT_NUMBER_OF_PAGES_TO_SHOW } from "../appSettings";
    import "../styles/components/pager.css";

    export let maxPages;
    export let currentPage;
    export let pageQueryString;

    const pagesToShow: number[] = [currentPage];

    const openPage = (pageNumber: number) => {
        window.open(`${pageQueryString}=${pageNumber}`, "_self");
    }

    for (let i = 1; i < 2*DEFAULT_NUMBER_OF_PAGES_TO_SHOW; i++) {
        if ((currentPage - i) > 0) {
            pagesToShow.push(currentPage - i)
        }
        if ((currentPage + i) <= maxPages) {
            pagesToShow.push(currentPage + i)
            console.log(pagesToShow);
        }
        if (pagesToShow.length > 10) {
            break;
        }
    }

    pagesToShow.sort((a, b) => {return a - b});
</script>

<div class="pagerContainer">
    {#if currentPage - 1 > 0}
    <button on:click={() => openPage(currentPage - 1)}>&lt</button>
    {/if}
    {#each pagesToShow as pageToShow}
        {#if pageToShow === currentPage}
            <button class="currentPage" on:click={() => openPage(pageToShow)}>{pageToShow}</button>
        {:else}
            <button on:click={() => openPage(pageToShow)}>{pageToShow}</button>
        {/if}
    {/each}
    {#if currentPage + 1 <= maxPages}
        <button on:click={() => openPage(currentPage + 1)}>&gt</button>
    {/if}
</div>