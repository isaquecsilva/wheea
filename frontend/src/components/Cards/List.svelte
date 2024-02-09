<script type="text/javascript">
	import { onMount } from 'svelte'

	export let options = [];
	export let coords = {}
	export let theme;

	onMount(() => {
		window.addEventListener('queryplaceresponse', event => {
			document.getElementById('list-container').style.height = '16vh'

			options = event.detail.data.reduce((accumulator, item) => {
				(accumulator.find(el => el == `${item.Name}, ${item.Country}`) == undefined) ? accumulator.push(`${item.Name}, ${item.Country}`) : null;
				return accumulator
			}, [])
		})
	})

	function queryWheather(cityname) {
		document.getElementById('list-container').style.height = '0'

		options = [];
		console.table({
			operation: 'api-call',
			param: cityname,
		})
	}
</script>

<div id="list-container" style="grid-row-start:{coords.rs};grid-row-end:{coords.re};grid-column-start:{coords.cs};grid-column-end:{coords.ce};
--list-theme: {theme == '#111' ? '#333' : '#eee'};">
	{#each options as option}
		<p on:click={(event) => queryWheather(event.target.innerText)}>{option}</p>
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