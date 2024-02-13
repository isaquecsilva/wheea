const themekit = {
	lightmode: '#fff',
	darkmode: '#111',
}

function getUserTheme() {
	return localStorage.getItem("apptheme")
}

function setUserTheme(theme) {
	localStorage.setItem("apptheme", theme)
}

function setComponentGridCoords(rs=1, re=1, cs=1, ce=1) {
	return {rs,	re,	cs,	ce}
}

export {
	themekit,
	getUserTheme,
	setUserTheme,
	setComponentGridCoords
}
