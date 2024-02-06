<script>
  import Search from './components/Fields/Search.svelte'
  import Theme from './components/Buttons/Theme.svelte'
  import WheatherCard from './components/Cards/WheatherCard.svelte'
  import { themekit, getUserTheme } from './theme.js'
  import wheather from './assets/images/illustrations/sun_rain.png'

  export let theme = getUserTheme() || themekit.darkmode
  $: document.body.style.backgroundColor = theme;
</script>

<main>
  <Theme id="theme" bind:theme />
  <WheatherCard id="card" bind:theme image={wheather}/>
  <Search id="search" bind:theme />
</main>

<style type="text/css">
  main {
    display: grid;
    height: 100%;
    width:  100%;
    grid-template-rows: 1fr 7fr 2fr;
  }

  #theme {
    grid-row-start: 0;
    grid-row-end: 1;
  }

  #card {
    grid-row-start: 1;
    grid-row-end: 2;
  }

  #search {
    grid-row-start: 2;
    grid-row-end: 3;
    height: 100vh;
    background: red;
  }
</style>