<script>
	export let theme;
	export let image = '';
	export let coords = {}
</script>

<div id="maincontainer" style="--wheather-card-theme: {theme};grid-row-start:{coords.rs};grid-row-end:{coords.re};grid-column-start:{coords.cs};grid-column-end:{coords.ce};">
	{#if image == ''}
		<span>No Wheather Informations Available</span>
	{:else}
		<div>
			<div class="elements">
				<h6>Wind: 18 km/h</h6>
				<h6>Precipitation: 16%</h6>
			</div>

			<img src={image} width="160" height="160" alt="wheather-image" />
			<div id="inner">
				<h4>Tokyo, Japan</h4>
				<div id="info">
					<div class="elements">
						<h6>Min: 18º</h6>
						<h6>Current: 32º</h6>
						<h6>Max: 33º</h6>
					</div>			
				</div>
			</div>
		</div>
	{/if}
</div>

<style type="text/css">
	#maincontainer {
		align-self: center;
		background-color: var(--wheather-card-theme);
		position: relative;
		height: 90%;
		width: 60vw;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin-left: 50%;
		transform: translateX(-50%);
		border: 1px solid dimgray;
		border-radius: 16px;
	}

	#maincontainer > div {
		display: flex;
		flex-direction: column;
		align-items: center;
		position: relative;
		width: 100%;
	}

	#maincontainer #inner {
		position: relative;
		width: 100%;
	}

	#maincontainer #inner > div {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		position: relative;
		width: 100%;
	}

	#info {
		display: flex;
		flex-direction: column;
		align-items: center;
		position: relative;
		width: 100vw;
	}

	.elements {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;
		justify-content: space-around;
		position: relative;
		width: 100%;
		background-color: none;
	}

	span {
		color: dimgray;
	}

	img {
		position: relative;
		width: 80px;
		height: 80px;
		max-width: 120px;
		max-height: 120px;
		margin-top: 12px;
		margin-bottom: 2%;
	}

	h4 {
		margin: 10px;
	}
</style>