import { WeatherQueryBind } from '../../wailsjs/go/main/WheeaApp.js'
import sweetalert2 from 'sweetalert2'

export function makeWeatherQuery({Name: name, Country: country, Latitude, Longitude}) {
    WeatherQueryBind(Latitude, Longitude)
		.then(data => {
            console.log(data);
			data.Result.Today.Name = name
			data.Result.Today.Country = country
			// Emitting an event to App.svelte for re-render its components
			// based on our acquired data.
			let event = new CustomEvent('weatherqueryresponse', {
				detail: {
					data,
				}
			})

			dispatchEvent(event);
        })
		.catch(error => {
			sweetalert2.fire({
				icon: 'error',
				title: 'Error Fetching Weather Infos',
				text: error.message,
			})
		})
}