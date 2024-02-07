const waitTime = 3 * 1000;
var apiCallCountdown = null;

export function getCityInfo(cityname) {
	if (apiCallCountdown != null) {
		clearTimeout(apiCallCountdown);
	}

	apiCallCountdown = setTimeout(() => console.log(`calling api: ${cityname}`), 
		waitTime)
}