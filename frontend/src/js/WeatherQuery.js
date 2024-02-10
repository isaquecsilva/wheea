import { WeatherQueryBind } from '../../wailsjs/go/main/WheeaApp.js'

export function makeWeatherQuery(lat, lon) {
    return WeatherQueryBind(lat, lon)
}