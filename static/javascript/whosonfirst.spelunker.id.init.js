window.addEventListener("load", function load(event){

    // START OF wrap me in a webcomponent

    whosonfirst.spelunker.yesnofix.enabled(false);
    
    try {

	var map_el = document.querySelector("#map");
	var svg_el = document.querySelector("#map-svg");	
	var wof_id = map_el.getAttribute("data-wof-id");

	whosonfirst.spelunker.feature.fetch(wof_id).then((f) => {

	    map_el.style.display = "block";
	    
	    const map = L.map(map_el);

	    if (f.geometry.type == "Point"){

		var coords = f.geometry.coordinates;
		
		var pt = [ coords[1], coords[0] ];
		var zm = Math.max(12, f.properties["mz:min_zoom"]);
		map.setView(pt, zm);
		
	    } else {
		var bounds = whosonfirst.spelunker.geojson.derive_bounds(f);
		map.fitBounds(bounds);
	    }

	    var tile_url = "https://static.sfomuseum.org/pmtiles/sfomuseum_v3/{z}/{x}/{y}.mvt?key=xxx";
	    var layer = protomapsL.leafletLayer({url: tile_url});
	    layer.addTo(map);

	    // http://localhost:8080/id/1259472055
	    if (f.geometry.type == "Point"){

		var pt_handler = whosonfirst.spelunker.leaflet.handlers.point();
		var lbl_style = whosonfirst.spelunker.leaflet.styles.label_centroid();	    		
		whosonfirst.spelunker.leaflet.draw_point(map, f, lbl_style, pt_handler);
		
		return;
	    }
	    
	    var f_style = whosonfirst.spelunker.leaflet.styles.consensus_polygon();
	    whosonfirst.spelunker.leaflet.draw_poly(map, f, f_style);

	    var props = f.properties;

	    var lbl_centroid = [ props["lbl:longitude"], props["lbl:latitude" ] ];
	    var math_centroid = [ props["geom:longitude"], props["geom:latitude" ] ];	    

	    var lbl_f = { "type": "Feature", "properties": { "lflt:label_text": "label centroid" }, "geometry": { "type": "Point", "coordinates": lbl_centroid }};
	    var math_f = { "type": "Feature", "properties": { "lflt:label_text": "math centroid" }, "geometry": { "type": "Point", "coordinates": math_centroid }};	    

	    var lbl_style = whosonfirst.spelunker.leaflet.styles.label_centroid();
	    var math_style = whosonfirst.spelunker.leaflet.styles.math_centroid();	    

	    var pt_handler = whosonfirst.spelunker.leaflet.handlers.point();
	    whosonfirst.spelunker.leaflet.draw_point(map, math_f, math_style, pt_handler);
	    whosonfirst.spelunker.leaflet.draw_point(map, lbl_f, lbl_style, pt_handler);
	    
	}).catch((err) => {
	    console.log("Failed to initialize map", err);
	    throw(err);
	});
	
    } catch (err) {
	console.log("Failed to initialize map", err);
	svg_el.style.display = "block";	    	
    }
    
    // END OF wrap me in a webcomponent    
    
    // START OF wrap me in a webcomponent

    var pretty;
    
    try {
	var el = document.querySelector("#whosonfirst-properties");
	var raw = el.innerText;
	var props = JSON.parse(raw);
	pretty = whosonfirst.spelunker.properties.render(props);	
    } catch(err) {
	console.log("Failed to render properties", err);
    }

    try {
        var wrapper = document.querySelector("#props-wrapper");
        wrapper.appendChild(pretty);
	
        var raw = wrapper.children[0];
        raw.style.display = "none";

        //wrapper.replaceChild(pretty, raw);
	
        var toggle = document.querySelector("#props-toggle");
        toggle.style.display = "inline-block";
	
        var toggle_raw = document.querySelector("#props-toggle-raw");
        toggle_raw.style.display = "block";
	
        toggle_raw.onclick = function(){	    
	    raw.style.display = "block";
            pretty.style.display = "none";	    
            toggle_raw.style.display = "none";
            toggle_pretty.style.display = "block";
        };
	
        var toggle_pretty = document.querySelector("#props-toggle-pretty");

	toggle_pretty.onclick = function(){	    
            raw.style.display = "none";
	    pretty.style.display = "block";	    
            toggle_raw.style.display = "block";
	    toggle_pretty.style.display = "none";
        };
	
    } catch(err){
	console.log("Failed to install pretty properties", err);
    }

    // END OF wrap me in a webcomponent

    whosonfirst.spelunker.namify.namify_selector(".props-uoc");
    whosonfirst.spelunker.namify.namify_selector(".wof-namify");    
});
