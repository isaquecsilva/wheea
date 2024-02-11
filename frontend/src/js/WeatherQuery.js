import { WeatherQueryBind } from '../../wailsjs/go/main/WheeaApp.js'
import sweetalert2 from 'sweetalert2'

export function makeWeatherQuery({Name, Latitude, Longitude}) {
    WeatherQueryBind(Latitude, Longitude)
		.then(data => {
            console.log(data);
        })
		.catch(error => {
			sweetalert2.fire({
				icon: 'error',
				title: 'Error Fetching Weather Infos',
				text: error.message,
			})
		})
}