window.addEventListener("load", function load(event){

    // Because we still don't have a place to send/store these data
    whosonfirst.spelunker.yesnofix.enabled(false);
    
    // START OF wrap me in a webcomponent

    const bbox_pane_name = "bbox";
    const bbox_pane_zindex = 1000;
    
    const parent_pane_name = "parent";
    const parent_pane_zindex = 2000;
    
    const poly_pane_name = "polygon";
    const poly_pane_zindex = 3000;
    
    const centroids_pane_name = "centroids";
    const centroids_pane_zindex = 4000;
    
    try {

	var map_el = document.querySelector("#map");
	var svg_el = document.querySelector("#map-svg");	
	var wof_id = map_el.getAttribute("data-wof-id");

	whosonfirst.spelunker.feature.fetch(wof_id).then((f) => {

	    map_el.style.display = "block";
	    
	    const map = L.map(map_el);

	    var bbox_pane = map.createPane(bbox_pane_name);
	    bbox_pane.style.zIndex = bbox_pane_zindex;
	    
	    var parent_pane = map.createPane(parent_pane_name);
	    parent_pane.style.zIndex = parent_pane_zindex;
	    
	    var poly_pane = map.createPane(poly_pane_name);
	    poly_pane.style.zIndex = poly_pane_zindex;
	    
	    var centroids_pane = map.createPane(centroids_pane_name);
	    centroids_pane.style.zIndex = centroids_pane_zindex;
	    
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

		var pt_handler_layer_args = {
		    pane: centroids_pane_name,
		};
		
		var pt_handler = whosonfirst.spelunker.leaflet.handlers.point(pt_handler_layer_args);
		var lbl_style = whosonfirst.spelunker.leaflet.styles.label_centroid();

		var layer_args = {
		    style: lbl_style,
		    pointToLayer: pt_handler,
		    pane: centroids_pane_name,
		}
		
		whosonfirst.spelunker.leaflet.draw_point(map, f, layer_args);
		
		return;
	    }

	    var bbox_style = whosonfirst.spelunker.leaflet.styles.bbox();

	    var bbox_layer_args = {
		style: bbox_style,
		pane: bbox_pane_name,
	    }
	    
	    whosonfirst.spelunker.leaflet.draw_bbox(map, f, bbox_layer_args);
	    
	    var poly_style = whosonfirst.spelunker.leaflet.styles.consensus_polygon();
	    
	    var poly_layer_args = {
		style: poly_style,
		pane: poly_pane_name,
	    };
	    
	    whosonfirst.spelunker.leaflet.draw_poly(map, f, poly_layer_args);

	    var props = f.properties;

	    var lbl_centroid = [ props["lbl:longitude"], props["lbl:latitude" ] ];
	    var math_centroid = [ props["geom:longitude"], props["geom:latitude" ] ];	    

	    var lbl_f = {
		"type": "Feature",
		"properties": { "lflt:label_text": "label centroid" },
		"geometry": { "type": "Point", "coordinates": lbl_centroid }
	    };
	    
	    var math_f = {
		"type": "Feature",
		"properties": { "lflt:label_text": "math centroid" },
		"geometry": { "type": "Point", "coordinates": math_centroid }
	    };	    

	    var lbl_style = whosonfirst.spelunker.leaflet.styles.label_centroid();
	    var math_style = whosonfirst.spelunker.leaflet.styles.math_centroid();

	    var pt_handler_layer_args = {
		pane: centroids_pane_name,
	    };
	    
	    var pt_handler = whosonfirst.spelunker.leaflet.handlers.point(pt_handler_layer_args);

	    var math_layer_args = {
		style: math_style,
		pointToLayer: pt_handler,
		pane: centroids_pane_name,
	    };

	    var lbl_layer_args = {
		style: lbl_style,
		pointToLayer: pt_handler,
		pane: centroids_pane_name,
	    };
	    
	    whosonfirst.spelunker.leaflet.draw_point(map, math_f, math_layer_args);
	    whosonfirst.spelunker.leaflet.draw_point(map, lbl_f, lbl_layer_args);

	    // Draw parent here...

	    var parent_id = f.properties["wof:parent_id"];
	    console.log("Fetch parent", parent_id);
	    
	    whosonfirst.spelunker.feature.fetch(parent_id).then((parent_f) => {

		if (! parent_f.geometry.type.endsWith("Polygon")){
		    return;
		}

		var parent_style = whosonfirst.spelunker.leaflet.styles.parent_polygon();
		
		var parent_layer_args = {
		    style: parent_style,
		    pane: parent_pane_name,
		};
		
		whosonfirst.spelunker.leaflet.draw_poly(map, parent_f, parent_layer_args);
		
	    }).catch((err) => {
		console.log("Failed to fetch parent record", err);
	    })
	    
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

    var props;
    var pretty;
    
    try {
	var el = document.querySelector("#whosonfirst-properties");
	var raw = el.innerText;
	props = JSON.parse(raw);
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

    var map_wrapper = document.querySelector("#map-wrapper");

    var is_ceased = false;
    var is_deprecated = false;    
    var is_superseded = false;
    var is_superseding = false;
    
    if ((props["edtf:deprecated"]) && (props["edtf:deprecated"] != "")){
	console.log("WTF", props["edtf:deprecated"]);
	is_deprecated = true;
    }

    if ((props["edtf:cessation"] != "") && (props["edtf:cessation"] != "uuuu")){
	is_ceased = true;
    }

    if (is_deprecated){
	var span = document.createElement("span");
	span.setAttribute("class", "hey-look");
	span.setAttribute("style", "color:red;");
	span.appendChild(document.createTextNode("This record is deprecated. "));
	map_wrapper.appendChild(span);

	document.querySelector("#writefield-link").style.display = "none";
    } else if (is_ceased){
	var el = document.createTextNode("This record is ceased. ");
	map_wrapper.appendChild(el);
    } else {

    }

    var count_supersedes = props["wof:supersedes"].length;
    var count_superseded_by = props["wof:superseded_by"].length;    

    if (count_supersedes > 0){
	var span = document.createElement("span");
	span.appendChild(document.createTextNode("This record "));

	var sup_span = document.createElement("span");
	sup_span.setAttribute("class", "hey-look");
	sup_span.appendChild(document.createTextNode("supersedes"));

	span.appendChild(sup_span);
	
	for (var i=0; i < count_supersedes; i++){
	    var a = document.createElement("a");
	    a.setAttribute("href", "#");
	    a.setAttribute("class", "wof-namify");
	    a.setAttribute("data-wof-id", props["wof:supersedes"][i]);
	    a.appendChild(document.createTextNode(props["wof:supersedes"]));

	    var c = document.createElement("code");
	    c.appendChild(document.createTextNode(props["wof:supersedes"][i]));
	    
	    span.appendChild(a);
	    span.appendChild(c);	    

	    if (i < (count_supersedes-1)){
		span.appendChild(document.createTextNode(", "));
	    }
	}

	span.appendChild(document.createTextNode(". "));
	map_wrapper.appendChild(span);
    }

    if (count_superseded_by > 0){
	var span = document.createElement("span");
	span.appendChild(document.createTextNode("This record has been "));

	var sup_span = document.createElement("span");
	sup_span.setAttribute("class", "hey-look");
	sup_span.appendChild(document.createTextNode("superseded"));

	span.appendChild(sup_span);
	span.appendChild(document.createTextNode(" by "));

	for (var i=0; i < count_superseded_by; i++){
	    var a = document.createElement("a");
	    a.setAttribute("href", "#");
	    a.setAttribute("class", "wof-namify");
	    a.setAttribute("data-wof-id", props["wof:superseded_by"][i]);
	    a.appendChild(document.createTextNode(props["wof:superseded_by"]));

	    var c = document.createElement("code");
	    c.appendChild(document.createTextNode(props["wof:superseded_by"][i]));
	    
	    span.appendChild(a);
	    // span.appendChild(document.createTextNode("("));
	    span.appendChild(c);
	    // span.appendChild(document.createTextNode(")"));
	    
	    if (i < (count_supersedes-1)){
		span.appendChild(document.createTextNode(", "));
	    }
	}

	span.appendChild(document.createTextNode(". "));	
	map_wrapper.appendChild(span);
    }
    
    //
    
    whosonfirst.spelunker.namify.namify_selector(".props-uoc");
    whosonfirst.spelunker.namify.namify_selector(".wof-namify");
    whosonfirst.spelunker.namify.namify_selector(".yesnofix-uoc");
    
});
