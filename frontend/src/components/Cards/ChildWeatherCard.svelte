<script>
    import { weatherCodeToImage } from '../../js/utils/WeatherCardUtils'
    export let weatherData = {
        MaxTemperature: 0,
        MinTemperature: 0,
        WeatherType: 0,
    }

</script>

<div id="container">
    <img src="{weatherCodeToImage(weatherData.WeatherType)}" alt="">
    <div>
        <h6>Min: {weatherData.MinTemperature}º</h6>
        <h6>Max: {weatherData.MaxTemperature}º</h6>
    </div>
</div>

<style>
    #container {
        border: 1px solid dimgray;
        border-radius: 8px;
        position: relative;
        width:  40%;
        height: 80%;
        margin: 0 6px;
        padding: 4px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    img {
        min-height: 16px;
        max-height: 32px;
        min-width:  16px;
        max-width:  32px;

        margin-bottom: 4px;
    }

    #container div {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
    }

    h6 {
        margin: 0;
    }
</style>