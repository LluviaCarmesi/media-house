<script>
    import "../styles/components/filePicker.css";
    export let inputID;
    export let fieldLabel;
    export let onFileChange;
</script>

<div class="filePickerContainer">
    <span>{fieldLabel}</span>
    <input id={inputID} type="file" on:change={onFileChange} />
</div>
