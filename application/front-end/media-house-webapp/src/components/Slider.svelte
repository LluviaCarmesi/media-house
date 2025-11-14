<script>
    import "../styles/components/slider.css";
    export let inputID;
    export let fieldLabel;
    export let onSliderChange;
    export let onSliderInput;
    export let currentValue;
    export let startLabel;
    export let endLabel;
    export let minimum;
    export let maximum;
</script>

<div class="sliderContainer">
    <span>{fieldLabel}</span>
    <span>{startLabel}</span>
    <input
        type="range"
        id={inputID}
        min={minimum}
        max={maximum}
        value={currentValue}
        on:input={onSliderInput}
        on:change={onSliderChange}
    />
    <span>{endLabel}</span>
</div>
