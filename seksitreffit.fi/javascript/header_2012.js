// JavaScript Document ilman pop


function toggle_visible(id)
{
    var elem = document.getElementById(id);
    if (elem.style.display == 'none') {
        elem.style.display = 'block';
    } else {
        elem.style.display = 'none';
    }
}


function toggle_vis_class(id)
{
/*
	for (sulje=1; sulje<=1; sulje++)
	{
		if (id != sulje) {
			var elem = document.getElementById(sulje);
			var isHidden = new RegExp('\\bhidden\\b').test(elem.className);
			
			if( !isHidden ) {
				elem.className = elem.className.replace('visible', 'hidden');
			}
		}
	}
*/
	var elem = document.getElementById(id);
    var isHidden = new RegExp('\\bhidden\\b').test(elem.className);
	
	if( isHidden ) {
        elem.className = elem.className.replace('hidden', 'visible');
    } else {
        elem.className = elem.className.replace('visible', 'hidden');
    }
}


function show_element(id)
{
    var elem = document.getElementById(id);
    elem.style.display = 'block';
}




var Toggle = { display: Element.toggle };



