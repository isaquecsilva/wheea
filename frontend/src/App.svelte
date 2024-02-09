<script>
  import Search from './components/Fields/Search.svelte'
  import Theme from './components/Buttons/Theme.svelte'
  import WheatherCard from './components/Cards/WheatherCard.svelte'
  import List from './components/Cards/List.svelte'

  import { themekit, getUserTheme, setComponentGridCoords } from './theme.js'

  export let theme = getUserTheme() || themekit.darkmode
  $: document.body.style.backgroundColor = theme;
</script>

<main>
  <Theme coords={setComponentGridCoords(1, 1, 1, 4)} bind:theme />
  <WheatherCard coords={setComponentGridCoords(2, 2, 1, 4)} bind:theme image="" />
  <List coords={setComponentGridCoords(3, 3, 2, 3)} bind:theme />
  <Search coords={setComponentGridCoords(4, 4, 1, 4)} bind:theme />
</main>

<style type="text/css">
  main {
    display: grid;
    height: 100%;
    width:  100%;
    grid-template-rows: 1fr 6fr 1fr 2fr;
    grid-template-columns: 1fr 4fr 1fr;
  }
</style>