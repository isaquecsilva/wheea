<script>
  import Search from './components/Fields/Search.svelte'
  import Theme from './components/Buttons/Theme.svelte'
  import WheatherCard from './components/Cards/WheatherCard.svelte'
  import ChildWeatherCard from './components/Cards/ChildWeatherCard.svelte'
  import { themekit, getUserTheme, setComponentGridCoords } from './theme.js'
  import { onMount } from 'svelte'
  import { weatherCodeToImage } from './js/utils/WeatherCardUtils';
  import { makeWeatherQuery } from './js/WeatherQuery'

  export let theme = getUserTheme() || themekit.darkmode
  $: document.body.style.backgroundColor = theme;

  var weatherdata = null;
  var weatherNextDaysData = [];
  let image = ''

  onMount(() => {
    window.addEventListener('weatherqueryresponse', function(event) {
      weatherdata = event.detail.data.Result.Today;
      weatherNextDaysData = event.detail.data.Result.NextDays;
      image = weatherCodeToImage(weatherdata.WeatherType);
    })

    var lastFetchedCity = localStorage.getItem("lastcity")
    if (lastFetchedCity) {
      let lastcity = JSON.parse(lastFetchedCity);
      console.log(lastcity)
      makeWeatherQuery(lastcity)
    }
  })  
</script>

<main>
  <Theme coords={setComponentGridCoords(1, 1, 1, 4)} bind:theme />
  <WheatherCard coords={setComponentGridCoords(2, 2, 2, 3)} bind:theme {image} {weatherdata} />
  <div id="nextdays">
    {#each weatherNextDaysData as weatherData}
      <ChildWeatherCard {weatherData} />
    {/each}
  </div>
  <Search coords={setComponentGridCoords(4, 4, 1, 4)} bind:theme />
</main>

<style type="text/css">
  main {
    display: grid;
    height: 100%;
    width:  100%;
    grid-template-rows: 1fr 4fr 3fr 2fr;
    grid-template-columns: 1fr 4fr 1fr;
  }

  #nextdays {
    background-color: none;
    grid-row-start: 3;
    grid-row-end: 3;
    grid-column-start: 1;
    grid-column-end: 4;
    position: relative;
    width: 100vw;
    display: flex;
    flex-wrap: nowrap;
    flex-direction: row;
    overflow-x: hidden;
    align-items: center;
    justify-content: center;
  }
</style>