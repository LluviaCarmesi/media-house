<script lang="ts">
    import "../../../styles/common.css";
    import Navigation from "../../../components/Navigation.svelte";
</script>

<Navigation />

<div>
    <div class="textContainer">
        <h2 class="mainHeading">Edit the Video</h2>
    </div>
    <div class="textContainer">
        <span class="description">Edit the video using the form below! </span>
    </div>
    <!--
    {#if !!progress}
        <div class="textContainer">
            <span class="warning"
                >Editing the video! Don't leave! {progress}% complete</span
            >
        </div>
    {/if}
    {#if !!errorMessage}
        <div class="textContainer">
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="textContainer">
            <span class="success">Video was added</span>
        </div>
    {/if}
    <div class="form">
        <div class="centerContainer">
            <TextField
                inputID="Title"
                currentValue={videoItem.Title}
                fieldLabel="Title"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <Dropdown
                inputID="Type"
                currentValue={videoItem.Type}
                fieldLabel="Type"
                onDropdownChange={handleDropdownChange}
                dropdownOptions={TYPES_OPTIONS}
            />
        </div>
        {#if videoItem.Type == TYPES.SHOW.value}
            <div class="centerContainer">
                <TextField
                    inputID="Episode"
                    currentValue={videoItem.Episode}
                    fieldLabel="Episode Number"
                    onChangeTextField={handleTextChange}
                />
            </div>
            <div class="centerContainer">
                <Dropdown
                    inputID="ShowID"
                    currentValue={videoItem.ShowID}
                    fieldLabel="Show"
                    onDropdownChange={handleDropdownChange}
                    dropdownOptions={[]}
                />
            </div>
        {/if}
        <div class="centerContainer">
            <TextField
                inputID="Language"
                currentValue={videoItem.Language}
                fieldLabel="Language"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="Duration"
                currentValue={videoItem.Duration}
                fieldLabel="Duration"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <TextField
                inputID="TagsString"
                currentValue={videoItem.TagsString}
                fieldLabel="Tags"
                onChangeTextField={handleTextChange}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="PreviewFile"
                fieldLabel="Attach Preview"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="centerContainer">
            <FilePicker
                inputID="VideoFile"
                fieldLabel="Attach Video"
                onFileChange={handleFileUpload}
            />
        </div>
        <div class="endContainer">
            <button on:click={submitVideo}>Submit</button>
        </div>
    </div>
    -->
</div>