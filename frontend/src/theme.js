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

export {
	themekit,
	getUserTheme,
	setUserTheme
}
