
$(document).ready(function(){
	$('a[href*="/go101.org/"]').each(function () {
		let p = "/go101.org"
		let h = $(this).attr("href")
		let i = h.indexOf(p)
		if (i >= 0) {
			$(this).attr("href", h.substr(i+p.length))
		}
	});

	// theme

	var remove = function(id) {
		var node = document.getElementById(id)
		node.parentNode.removeChild(node)
	}

	var loadCSS = function(id, cssPath) {
		remove(id)
		var head = document.getElementsByTagName('head')[0]
		var link = document.createElement('link')
		link.id = id
		link.rel = 'stylesheet'
		link.type = 'text/css'
		link.href = cssPath
		link.media = 'all'
		head.appendChild(link)
	}

	var loadJS = function(id, jsPath) {
		remove(id)
		var head  = document.getElementsByTagName('head')[0]
		var script = document.createElement('script')
		script.id = id
		script.setAttribute("type","text/javascript")
		script.setAttribute("src", jsPath)
		head.appendChild(script);
	}

	var swithTheme = function(targetTheme, manually) {
		if (theme == targetTheme) {
			return
		}

		theme = targetTheme
		if (theme == "light") {
			loadCSS('css-bs', '/static/bootstrap/v4.5.0/css/bootstrap.min.css')
			loadCSS('css-prism', '/static/prism/2020-08-03-light/prism.css')
			loadJS('js-prism', '/static/prism/2020-08-03-light/prism.js')
			loadCSS('css-go101', '/static/go101/css/v993-light.css')
		} else if (theme == "dark") {
			loadCSS('css-bs', '/static/bootstrap/v4.0.3-dark/css/bootstrap.min.css')
			loadCSS('css-prism', '/static/prism/2020-08-03-dark/prism.css')
			loadJS('js-prism', '/static/prism/2020-08-03-dark/prism.js')
			loadCSS('css-go101', '/static/go101/css/v993-dark.css')
		}

		if (manually) {
			document.cookie = "theme=" + theme  + "; path=/; SameSite=None; Secure"
		}
	}

	var cookieTheme = theme
	if (cookieTheme == "") {
		document.cookie.split(';').forEach(function(el) {
			let [key, value] = el.split('=')
			if (key.trim() == "theme") {
				cookieTheme = value
			}
		})
		if (cookieTheme == "") {
			const mediaQueryList = window.matchMedia("(prefers-color-scheme: light)")
			if (mediaQueryList.matches) {
				cookieTheme = "light"
			}
		}
		if (cookieTheme == "") {
			cookieTheme = "dark"
		}
	}
	//console.log("=== theme: ", theme)
	//console.log("=== document.cookie: ", document.cookie)
	//console.log("=== cookieTheme: ", cookieTheme)
	swithTheme(cookieTheme, false)

	var themeSwitch = document.getElementById('theme-switch')
	if (themeSwitch == null) {
		return
	}
	themeSwitch.style.cursor = 'pointer'
	themeSwitch.style.color = ""
	themeSwitch.classList.add("active")
	themeSwitch.addEventListener("click", function() {
		if (theme != "dark") {
			swithTheme("dark", true)
		} else {
			swithTheme("light", true)
		}
	}); 
});


