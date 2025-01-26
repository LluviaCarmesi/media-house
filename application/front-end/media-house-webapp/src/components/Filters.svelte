<script lang="ts">
    import "../styles/components/filters.css";
    export let onChangeTextbox;
    export let onChangeTagsDropdown;
    export let onEnterTextbox: any = null;
    export let label;
    export let tagsDropdownOptions;
    export let currentSearchValue;
    import { onMount } from "svelte";

    function onEnter(event: any) {
        if (event.keyCode === 13) {
            onEnterTextbox();
        }
    }

    function onEnterButton(event: any) {
        onEnterTextbox();
    }

    onMount(() => {
        document
            .getElementById("searchBar")
            ?.addEventListener("keyup", onEnter);
    });
</script>

<div class="filterContainer">
    <span>{label}&nbsp;</span>
    <input
        id="searchBar"
        value={currentSearchValue}
        type="text"
        on:keyup={onChangeTextbox}
    />
    {#if onEnterTextbox}
        <button on:click={onEnterButton}>Search</button>
    {/if}
    {#if tagsDropdownOptions && tagsDropdownOptions.length > 0}
        <select on:change={onChangeTagsDropdown}>
            <option value="Any">Any</option>
            {#each tagsDropdownOptions as tagDropdownOption}
                <option value={tagDropdownOption.value}>{tagDropdownOption.label}</option>
            {/each}
        </select>
    {/if}
</div>
