<script type="text/javascript">
	import { onMount } from 'svelte'
	import { makeWeatherQuery } from '../../js/WeatherQuery.js'
	import sweetalert2 from 'sweetalert2'

	export let options = [];
	export let coords = {}
	export let theme;

	onMount(() => {
		window.addEventListener('queryplaceresponse', event => {
			document.getElementById('list-container').style.height = '16vh'

			options = event.detail.data.reduce((accumulator, item) => {
				(accumulator.find(el => el == item) == undefined) ? accumulator.push(item) : null;
				return accumulator
			}, [])
		})
	})

	function queryWeather(event) {
		const queryPlaceObject = JSON.parse(atob(event.target.id))
		document.getElementById('list-container').style.height = '0'

		options = [];

		makeWeatherQuery(queryPlaceObject.Latitude, queryPlaceObject.Longitude)
			.then(data => console.log(data))
			.catch(error => {
				sweetalert2.fire({
					icon: 'error',
					title: 'Error Fetching Weather Infos',
					text: error.message,
				})
			})

		console.table({
			operation: 'api-call',
			param: queryPlaceObject.Name,
		})
	}
</script>

<div id="list-container" style="grid-row-start:{coords.rs};grid-row-end:{coords.re};grid-column-start:{coords.cs};grid-column-end:{coords.ce};--list-theme: {theme == '#111' ? '#333' : '#eee'};">
	{#each options as option}
		<p
		id={btoa(JSON.stringify(option))} 
		on:click={event => { queryWeather(event)}}>{`${option.Name}, ${option.Country}`}</p>
	{/each}
</div>

<style type="text/css">
	div {
		display: block;
		height: 0;
		background-color: var(--list-theme);
		border-radius: 6px 6px 0px 0px;
		transition: height 0.5s ease-in;
		padding: 0;
		overflow-y: auto;
		position: inherit;
	}

	p {
		background-color: none;
		margin: 0;
		padding: 8px;
		text-align: left;
		user-select: none;
		color: #696969;
		transition: 0.4s ease-out;
	}

	p:hover {
		cursor: pointer;
		background-color: #656565;
		color: white;
	}

	::-webkit-scrollbar {
		width: 6px;
		background-color: #555;
	}

	/* Track */
	::-webkit-scrollbar-track {
		background: #f1f1f1;
	}

	/* Handle */
	::-webkit-scrollbar-thumb {
		background: #888;
	}

	/* Handle on hover */
	::-webkit-scrollbar-thumb:hover {
		background: #555;
	}
</style>