<script lang="ts">
    import "../../../styles/pages/video.css";
    import "../../../styles/common.css";
    import Navigation from "../../../components/Navigation.svelte";
    export let data;
    import getVideoByID from "../../../services/getVideoByID";
    import { VIDEOS_SERVER_URI } from "../../../appSettings";
    import deleteVideoByID from "../../../services/deleteVideoByID";
    import Slider from "../../../components/Slider.svelte";
    import { onMount } from "svelte";
    import Checkbox from "../../../components/Checkbox.svelte";
    import formatSecondsToMinutesAndSeconds from "../../../utilities/formatSecondsToMinutes";

    let video = {
        ID: "",
        Title: "",
        Duration: "",
        Type: "",
        Tags: [""],
        PreviewPath: "",
        VideoPath: "",
    };
    let isLoading = true;
    const loadingMessage = "Loading...";

    let errorMessage = "";
    let videoInterval;
    let videoStopTime = new Date();

    let checkboxValues: any = {
        shouldRepeatVideo: false,
        shouldCloseVideoAtCertainTime: false,
    };
    const maxSliderValue = "1000";
    const maxSliderValueNumber = 1000;

    let sliderValues: any = {
        videoRepeatStartTime: "0",
        videoRepeatEndTime: maxSliderValue,
        closeVideoTime: "100",
    };

    let sliderValueLabels: any = {
        videoRepeatStartTimeLabel: "",
        videoRepeatEndTimeLabel: "",
        videoEndTimeLabel: "",
        closeVideoTimeLabel: "",
    };

    onMount(() => {
        videoInterval = setInterval(checkVideoInterval, 1000);
    });

    function onCheckboxChange(e: any) {
        checkboxValues[e.target.id] = e.target.checked;
        if (e.target.id === "shouldCloseVideoAtCertainTime") {
            const sliderValueNumber = parseInt(sliderValues.closeVideoTime);
            const currentTime = new Date();
            const secondsValue =
                ((sliderValueNumber * 60) / maxSliderValueNumber) * 60;
            sliderValueLabels.closeVideoTimeLabel =
                formatSecondsToMinutesAndSeconds(secondsValue);
            const millisecondsToAdd = secondsValue * 1000;
            videoStopTime = new Date(millisecondsToAdd + currentTime.getTime());
        }
    }

    function onSliderInput(e: any) {
        // Maybe create a hashmap of all possible time frames for the maxSliderValue
        // When changing slider input, check hashmap for value and update label
    }

    function onSliderChange(e: any) {
        sliderValues[e.target.id] = e.target.value;
        if (e.target.id === "closeVideoTime") {
            const sliderValueNumber = parseInt(sliderValues.closeVideoTime);
            const currentTime = new Date();
            const secondsValue =
                ((sliderValueNumber * 60) / maxSliderValueNumber) * 60;
            sliderValueLabels.closeVideoTimeLabel =
                formatSecondsToMinutesAndSeconds(secondsValue);
            const millisecondsToAdd = secondsValue * 1000;
            videoStopTime = new Date(millisecondsToAdd + currentTime.getTime());
        }
    }

    async function getVideoByIDResponse() {
        const videoResponse = await getVideoByID(data.id);
        if (videoResponse.isSuccessful) {
            video = videoResponse.video;
        } else {
            errorMessage = videoResponse.errorMessage;
        }
        isLoading = false;
    }

    async function deleteVideoByIDResponse() {
        errorMessage = "";
        const deleteVideoResponse = await deleteVideoByID(data.id);
        if (deleteVideoResponse.doesErrorExist) {
            errorMessage = deleteVideoResponse.errorMessage;
        } else {
            window.open("/", "_self");
        }
        isLoading = false;
    }

    function checkVideoInterval() {
        const videoElement: any = document.getElementById("videoElement");
        if (checkboxValues.shouldRepeatVideo) {
            if (!sliderValueLabels.videoEndTimeLabel) {
                sliderValueLabels.videoEndTimeLabel =
                    formatSecondsToMinutesAndSeconds(videoElement.duration);
                sliderValueLabels.videoRepeatStartTimeLabel = "0:00";
                sliderValueLabels.videoRepeatEndTimeLabel =
                    sliderValueLabels.videoEndTimeLabel;
            }

            const videoRepeatStartTimeNumber = parseInt(
                sliderValues.videoRepeatStartTime,
            );
            const videoRepeatEndTimeNumber = parseInt(
                sliderValues.videoRepeatEndTime,
            );
            const videoRepeatStartTimeDuration =
                (videoRepeatStartTimeNumber * videoElement.duration) /
                maxSliderValueNumber;
            const videoRepeatEndTimeDuration =
                (videoRepeatEndTimeNumber * videoElement.duration) /
                maxSliderValueNumber;

            sliderValueLabels.videoRepeatStartTimeLabel =
                formatSecondsToMinutesAndSeconds(videoRepeatStartTimeDuration);
            sliderValueLabels.videoRepeatEndTimeLabel =
                formatSecondsToMinutesAndSeconds(videoRepeatEndTimeDuration);

            if (videoElement.currentTime < videoRepeatStartTimeDuration) {
                videoElement.currentTime = videoRepeatStartTimeDuration;
                videoElement.play();
            }
            if (videoElement.currentTime >= videoRepeatEndTimeDuration) {
                videoElement.currentTime = videoRepeatStartTimeDuration;
                videoElement.play();
            }
        }

        if (checkboxValues.shouldCloseVideoAtCertainTime) {
            const currentTime = new Date();
            if (currentTime > videoStopTime) {
                videoElement.pause();
            }
        }
    }
    getVideoByIDResponse();
</script>

<Navigation />
<div class="container">
    {#if isLoading}
        <h1>{loadingMessage}</h1>
    {:else}
        <h1>{video.Title}</h1>
        <h1>{errorMessage}</h1>
        <video id="videoElement" controls>
            <source
                id="videoSource"
                src={`${VIDEOS_SERVER_URI}/${video.VideoPath}`}
                type="video/mp4"
            />
            Your browser doesn't support videos
            <track kind="captions" />
        </video>
        <Checkbox
            inputID="shouldRepeatVideo"
            fieldLabel="Should the video be repeated?"
            currentValue={checkboxValues.shouldRepeatVideo}
            {onCheckboxChange}
        />
        {#if checkboxValues.shouldRepeatVideo}
            <Slider
                inputID="videoRepeatStartTime"
                fieldLabel="Repeat Video at certain start time"
                {onSliderChange}
                {onSliderInput}
                currentValue={sliderValues.videoRepeatStartTime}
                startLabel={sliderValueLabels.videoRepeatStartTimeLabel}
                endLabel={sliderValueLabels.videoEndTimeLabel}
                minimum={"0"}
                maximum={maxSliderValue}
            />
            <Slider
                inputID="videoRepeatEndTime"
                fieldLabel="Repeat Video at certain end time"
                {onSliderChange}
                {onSliderInput}
                currentValue={sliderValues.videoRepeatEndTime}
                startLabel={sliderValueLabels.videoRepeatEndTimeLabel}
                endLabel={sliderValueLabels.videoEndTimeLabel}
                minimum={"0"}
                maximum={maxSliderValue}
            />
        {/if}
        <Checkbox
            inputID="shouldCloseVideoAtCertainTime"
            fieldLabel="Should close the video after certain time?"
            currentValue={checkboxValues.shouldCloseVideoAtCertainTime}
            {onCheckboxChange}
        />
        {#if checkboxValues.shouldCloseVideoAtCertainTime}
            <Slider
                inputID="closeVideoTime"
                fieldLabel="Stop video after certain amount of minutes (60 max)"
                {onSliderChange}
                {onSliderInput}
                currentValue={sliderValues.closeVideoTime}
                startLabel={sliderValueLabels.closeVideoTimeLabel}
                endLabel={"60"}
                minimum={"0"}
                maximum={maxSliderValue}
            />
        {/if}
        <!-- Have a dialog show up to confirm that the user wants to delete the video
        <div class="container">
            <button on:click={deleteVideoByIDResponse} class="deleteButton">Delete this video</button>
        </div> -->
    {/if}
</div>
