<script lang="ts">
    import { DEFAULT_NUMBER_OF_PAGES_TO_SHOW } from "../appSettings";
    import "../styles/components/pager.css";
    import { goto } from "$app/navigation";

    export let maxPages;
    export let currentPage;
    export let pageQueryString;

    const pagesToShow: number[] = [currentPage];

    console.log(currentPage);

    const openPage = (pageNumber: number) => {
        goto(`${pageQueryString}=${pageNumber}`);
    };

    for (let i = 1; i < 2 * DEFAULT_NUMBER_OF_PAGES_TO_SHOW; i++) {
        if (currentPage - i > 0) {
            pagesToShow.push(currentPage - i);
        }
        if (currentPage + i <= maxPages) {
            pagesToShow.push(currentPage + i);
        }
        if (pagesToShow.length > 10) {
            break;
        }
    }

    pagesToShow.sort((a, b) => {
        return a - b;
    });
</script>

<div class="pagerContainer">
    {#if currentPage - 1 > 0}
        <a
            on:click={() => openPage(currentPage - 1)}
            href={`${pageQueryString}=${currentPage - 1}`}>&lt</a
        >
    {/if}
    {#each pagesToShow as pageToShow}
        {#if pageToShow === currentPage}
            <a
                class="currentPage"
                on:click={() => openPage(pageToShow)}
                href={`${pageQueryString}=${pageToShow}`}>{pageToShow}</a
            >
        {:else}
            <a
                on:click={() => openPage(pageToShow)}
                href={`${pageQueryString}=${pageToShow}`}>{pageToShow}</a
            >
        {/if}
    {/each}
    {#if currentPage + 1 <= maxPages}
        <a
            on:click={() => openPage(currentPage + 1)}
            href={`${pageQueryString}=${currentPage + 1}`}>&gt</a
        >
    {/if}
</div>
