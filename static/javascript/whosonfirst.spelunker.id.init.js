window.addEventListener("load", function load(event){

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
        raw.setAttribute("style", "display:none");

        //wrapper.replaceChild(pretty, raw);
	
        var toggle = document.querySelector("#props-toggle");
        toggle.setAttribute("style", "display:block");
	
        var toggle_raw = document.querySelector("#props-toggle-raw");
        toggle_raw.setAttribute("style", "display:block");
	
        toggle_raw.onclick = function(){
	    
	    raw.setAttribute("style", "display:block");
            pretty.setAttribute("style", "display:none");
	    
            toggle_raw.setAttribute("style", "display:none");
            toggle_pretty.setAttribute("style", "display:block");
        };
	
        var toggle_pretty = document.querySelector("#props-toggle-pretty");

	toggle_pretty.onclick = function(){
	    
            raw.setAttribute("style", "display:none");
	    pretty.setAttribute("style", "display:block");
	    
            toggle_raw.setAttribute("style", "display:block");
	    toggle_pretty.setAttribute("style", "display:none");
        };
	
    } catch(err){
	console.log("Failed to install pretty properties", err);
    }

    // END OF wrap me in a webcomponent
});
