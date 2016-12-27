var url = window.location.href;
var xmlHttp = new XMLHttpRequest();
xmlHttp.open( "GET", "https://192.168.0.114:8080/stream/" + url.substring(url.lastIndexOf('?') + 1),  false );
xmlHttp.send( null );
