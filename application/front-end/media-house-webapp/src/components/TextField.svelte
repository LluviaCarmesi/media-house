<script>
    import "../styles/components/textField.css";
    export let fieldLabel;
    export let onChangeTextField;
    export let currentValue;
    export let inputID;
</script>

<div class="textFieldContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        value={currentValue}
        type="text"
        on:keyup={onChangeTextField}
        on:change={onChangeTextField}
    />
</div>
