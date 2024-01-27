<script>
	import { onMount } from 'svelte'
	export let theme;
</script>

<div style="--wheather-card-theme: {theme}"></div>

<style type="text/css">
	div {
		background-color: var(--wheather-card-theme);
		height: 80vh;
	}
</style>