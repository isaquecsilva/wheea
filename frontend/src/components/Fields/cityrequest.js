import { QueryWheather } from '../../../wailsjs/go/main/WheeaApp.js'
import sweetalert2 from 'sweetalert2'

const waitTime = 3 * 1000;
var apiCallCountdown = null;

export function getCityInfo(cityname) {
	if (apiCallCountdown != null) {
		clearTimeout(apiCallCountdown);
	}

	if (!cityname.length) {
		return
	}

	apiCallCountdown = setTimeout(async () => {
		let response = await QueryWheather(cityname)

		if (response.Error.length) {
			sweetalert2.fire({
				icon: "error",
				title: "Error",
				text: response.Error
			})
		} 
		else {
			var event = new CustomEvent('queryplaceresponse', { 
				detail: {
					data: response.Results
				}
			});
			window.dispatchEvent(event)
		}
	}, waitTime)
}