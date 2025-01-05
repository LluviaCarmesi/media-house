<script lang="ts">
    import "../../styles/common.css";
    import Navigation from "../../components/Navigation.svelte";
    import Filters from "../../components/Filters.svelte";
</script>

<Navigation />

<div>
    <div class="textContainer">
        <h2 class="mainHeading">Edit a Video</h2>
    </div>
    <div class="textContainer">
        <span class="description">
            Edit any video you see below!
        </span>
    </div>
    <Filters label="Search for a video"/>
</div>