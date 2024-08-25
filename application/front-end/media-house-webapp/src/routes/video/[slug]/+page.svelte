<script>
    import "../../../styles/pages/video.css";
    import Navigation from "../../../components/Navigation.svelte";
    import testingVideo from "../../../videos/Testing.mp4";
    export let data;
</script>

<Navigation />
<div class="container">
    <video src={testingVideo} controls></video>
</div>
