import { QueryPlaceBind, WeatherQueryBind } from '../../wailsjs/go/main/WheeaApp.js'
import { makeWeatherQuery } from './WeatherQuery.js';
import sweetalert2 from 'sweetalert2'

const waitTime = 3 * 1000;
var apiCallCountdown = null;

export function getCityInfo(cityname) {
	if (apiCallCountdown != null) {
		clearTimeout(apiCallCountdown);
	}

	if (!cityname.length) return

	apiCallCountdown = setTimeout(async () => {
		let response = await QueryPlaceBind(cityname)

		if (response.Error.length) {
			sweetalert2.fire({
				icon: "error",
				title: "Error",
				text: response.Error
			})
		} else {
			console.log(response);
			let { value } = await sweetalert2.fire({
				title: "Locations",
				input: "select",
				inputOptions: response.Results.reduce((accumulator, item, index) => {
					// Dont repeat location names
					if ( Object.values(accumulator).find(el => el == `${item.Name}, ${item.Country}`) == undefined ) {
						accumulator["city_"+index] = `${item.Name}, ${item.Country}`
					}

					return accumulator
				}, {}),
				inputPlaceholder: "Select a location",
				showCancelButton: true,
			})

			if (value) {
				let index = value.match(/\d+/gi)[0]
				let [ option ] = response.Results.filter((_, i) => i ==  index);
				makeWeatherQuery(option);
			}			
		}

	}, waitTime)
}