
$(document).ready(function(){
	$('a[href*="/go101.org/"]').each(function () {
		let p = "/go101.org"
		let h = $(this).attr("href")
		let i = h.indexOf(p)
		if (i >= 0) {
			$(this).attr("href", h.substr(i+p.length));
		}
	});
});

