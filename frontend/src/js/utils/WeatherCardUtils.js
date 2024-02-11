import { sun_rain_bolt, sunny, sun_cloudy, cloud_fog, cloudy_rain } from "../Assets";

export function weatherCodeToImage(weathercode) {
    console.log(`weather code: ${weathercode}`)
    switch (weathercode) {
        case 0||1:
            return sunny
        case 2:
            return sun_cloudy
        case 3:
            return cloud_fog
        case 45||48:
            return cloud_fog
        case 51||53||55||56||57||61||63||66||73||77||80||81||85:
            return cloudy_rain
        case 65||67||75||82||86||weathercode >= 95:
            return sun_rain_bolt
        default:
            return sunny

    }
}