<script>
  import Search from './components/Fields/Search.svelte'
  import Theme from './components/Buttons/Theme.svelte'
  import WheatherCard from './components/Cards/WheatherCard.svelte'
  import { themekit } from './theme.js'

  export let theme = themekit.theme;
  $: document.body.style.backgroundColor = theme;  
</script>

<header>
  <Theme bind:theme />
</header>

<main>
  <WheatherCard bind:theme />
</main>

<footer>
  <Search bind:theme />
</footer>