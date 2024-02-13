import { sun_rain_bolt, sunny, sun_cloudy, cloud_fog, cloudy_rain } from "../Assets";

export function weatherCodeToImage(weathercode) {
    console.log(`weather code: ${weathercode}`)
    switch (weathercode) {
        case 0:
        case 1:
            return sunny
        case 2:
            return sun_cloudy
        case 3:
            return cloud_fog
        case 45:
        case 48:
            return cloud_fog
        case 51:
        case 53:
        case 55:
        case 56:
        case 57:
        case 61:
        case 63:
        case 66:
        case 73:
        case 77:
        case 80:
        case 81:
        case 85:
            return cloudy_rain
        case 65:
        case 67:
        case 75:
        case 82:
        case 86:
        case weathercode >= 95:
            return sun_rain_bolt
        default:
            console.log('chegou aqui')
            return sunny
    }
}