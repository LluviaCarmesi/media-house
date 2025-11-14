<script>
    import "../styles/components/slider.css";
    export let inputID;
    export let fieldLabel;
    export let onSliderChange;
    export let currentValue;
    export let minimum;
    export let maximum;
</script>

<div class="sliderContainer">
    <span>{fieldLabel}</span>
    <input
        type="range"
        id={inputID}
        min={minimum}
        max={maximum}
        value={currentValue}
        on:change={onSliderChange}
    />
</div>
