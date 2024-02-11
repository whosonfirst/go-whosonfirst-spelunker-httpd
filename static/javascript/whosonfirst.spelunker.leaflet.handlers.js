var whosonfirst = whosonfirst || {};
whosonfirst.spelunker = whosonfirst.spelunker || {};
whosonfirst.spelunker.leaflet = whosonfirst.spelunker.leaflet || {};

whosonfirst.spelunker.leaflet.handlers = (function(){

	var self = {

		'point': function(layer_args){

			return function(feature, latlon){

				var m = L.circleMarker(latlon, layer_args);
				
				// https://github.com/Leaflet/Leaflet.label
				
				try {
					var props = feature['properties'];
					var label = props['lflt:label_text'];
					
					if (label){
						m.bindLabel(label, { noHide: false });
					}
				}
				
				catch (e){
					console.log("failed to bind label because " + e);
				}
				
				return m;
			};
		},
	};
	
	return self;
})();
