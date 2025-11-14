<script>
    import "../styles/components/checkbox.css";
    export let inputID;
    export let fieldLabel;
    export let onCheckboxChange;
    export let currentValue;
</script>

<div class="checkboxContainer">
    <span>{fieldLabel}</span>
    <input
        type="checkbox"
        id={inputID}
        value={currentValue}
        on:change={onCheckboxChange}
    />
</div>
