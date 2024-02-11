<script>
	import moon from '../../assets/images/themes/moon_theme_icon.png';
	import sun from '../../assets/images/themes/sun_theme_icon.png';
	import { themekit, getUserTheme, setUserTheme } from '../../theme.js';

	export let coords = {}
	export let theme;
	
	const size = { width: 20, height: 20, }
	let icon = getUserTheme() == themekit.lightmode ? sun : moon;

	// Change application current theme
	function changeTheme() {
		icon = icon == sun ? moon : sun;
		theme = (icon == moon) ? themekit.darkmode: themekit.lightmode;
		setUserTheme(theme)
	}
</script>
<div style="grid-row-start:{coords.rs};grid-row-end:{coords.re};grid-column-start:{coords.cs};grid-column-end:{coords.ce};">
	<button on:click={changeTheme}>
		<img src={icon} {...size} />
	</button>
</div>

<style>

	div {
		grid-row-start: 1;
	    grid-row-end: 2;
	    grid-column-start: 1;
	    grid-column-end: 4;
	}

	button {
		float: right;
		margin: 8px;
		background: transparent;
		border: none;
		outline: none;
	}
</style>