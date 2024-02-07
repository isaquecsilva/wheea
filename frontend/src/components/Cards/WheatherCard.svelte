<script>
	import { onMount } from 'svelte'

	export let theme;
	export let image = '';
</script>

<div style="--wheather-card-theme: {theme}">
	{#if image == ''}
		<span>No Wheather Informations Available</span>
	{:else}
		<img src={image} width="160" height="160" alt="wheather-image" />
	{/if}
</div>

<style type="text/css">
	div {
		background-color: var(--wheather-card-theme);
		position: relative;
		height: 100%;
		width: 100vw;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0;
	}

	span {
		color: dimgray;
	}
</style>