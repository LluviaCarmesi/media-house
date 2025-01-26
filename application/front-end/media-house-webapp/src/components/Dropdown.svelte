<script>
    import "../styles/components/dropdownField.css";
    export let inputID;
    export let fieldLabel;
    export let onDropdownChange;
    export let currentValue;
    export let dropdownOptions;
</script>

<div class="dropdownContainer">
    <span>{fieldLabel}</span>
    <select value={currentValue} id={inputID} on:change={onDropdownChange}>
        <option value="">Select</option>
        {#each dropdownOptions as dropdownOption}
            <option value={dropdownOption.value}>{dropdownOption.label}</option>
        {/each}
    </select>
</div>
