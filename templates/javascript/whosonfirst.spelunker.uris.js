{{ define "whosonfirst_spelunker_uris" -}}
var whosonfirst = whosonfirst || {};
whosonfirst.spelunker = whosonfirst.spelunker || {};

whosonfirst.spelunker.uris = (function(){

    var _table = {{ .Table }};

    var self = {
	
	abs_root_url: function(){
	    return "/";
	},
	
	table: function(){
	    return _table;
	},
    };

    return self;
})();
{{ end -}}
