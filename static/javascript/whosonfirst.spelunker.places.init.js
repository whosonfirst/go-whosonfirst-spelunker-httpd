window.addEventListener("load", function load(event){

    var places = document.querySelectorAll(".whosonfirst-places-list li");

    if (! places){
	console.log("No places");
	return;
    }
    
    var count_places = places.length;

    var coords = [];
    var names = [];
    
    for (var i=0; i < count_places; i++) {

	var el = places[i];
	var lat = parseFloat(el.getAttribute("data-latitude"));
	var lon = parseFloat(el.getAttribute("data-longitude"));	

	if ((! lat) || (!lon)){
	    console.log("Invalid coordinates", i, lat, lon);
	    continue;
	}

	var n = el.querySelector(".wof-place-name");

	if ((! n) || (n.innerText == "")){
	    console.log("Invalid name", i);
	    continue;
	}

	coords[i] = [ lon, lat ];
	names[i] = n.innerText;
    }

    console.log("Coords", coords);
    console.log("Names", names);
    
})();
