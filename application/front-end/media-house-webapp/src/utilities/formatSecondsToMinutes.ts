export default function formatSecondsToMinutesAndSeconds(totalSeconds: number) {
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    const formattedMinutes = minutes.toFixed(0);
    const formattedSeconds = seconds < 10 ? "0" + seconds.toFixed(0) : seconds.toFixed(0);
    
    return `${formattedMinutes}:${formattedSeconds}`;
}