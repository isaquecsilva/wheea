<script>
	import { getCityInfo } from '../../js/CityRequest.js';

	export let coords = {}
	export let theme;

	let cityname;
</script>

<div style="--search-theme: {theme};grid-row-start:{coords.rs};grid-row-end:{coords.re};grid-column-start:{coords.cs};grid-column-end:{coords.ce};">
	<span>City:</span>
	<input
		on:keyup={() => getCityInfo(cityname)}
		bind:value={cityname}
		placeholder="Tokyo"
	/>
</div>

<style>
	* {
		font-family: 'Roboto', 'Helvetica', sans-serif;
		font-weight: 600;
	}

	div {
		background-color: var(--search-theme);
		height: 20vh;
		border-radius: 12px 12px 0px 0px;
		box-shadow: 0px 1px 8px #333;
		display: grid;
		grid-template-rows: 1fr 2fr;
		grid-template-columns: 1fr 3fr 1fr;
		z-index: 1;
	}

	span {
		grid-row-start: 1;
		grid-column-start: 2;
		font-size: 12px;
		color: grey;
		text-align: left;
		align-self: flex-end;
		padding-bottom: 4px;
		padding-left: 6px;
	}

	input {
		grid-column-start: 2;
		grid-row-start: 2;
		grid-row-end: 2;
		position: relative;
		font-size: 16px;
		font-weight: 300;
		height: 2em;
		width: 66vw;
		border-radius: 6px;
		border: 1px solid slategray;
		outline: none;
		padding: 0px 12px;
		color: #444;
	}

	input::placeholder {
		color: #ccc;
	}
</style>