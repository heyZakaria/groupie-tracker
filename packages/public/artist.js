// Leaflet has native support for raster maps, So you can create a map with a few commands only!

// The Leaflet map Object
var map = L.map('my-map').setView([48.1500327, 11.5753989], 10);

// Get your own API Key on https://myprojects.geoapify.com
var myAPIKey = "ab95d9aafa1d45449323898be803875e";

// Retina displays require different mat tiles quality
var isRetina = L.Browser.retina;

var baseUrl = "https://maps.geoapify.com/v1/tile/osm-bright/{z}/{x}/{y}.png?apiKey=ab95d9aafa1d45449323898be803875e";
var retinaUrl = "https://maps.geoapify.com/v1/tile/osm-bright/{z}/{x}/{y}@2x.png?apiKey=ab95d9aafa1d45449323898be803875e";

// Add map tiles layer. Set 20 as the maximal zoom and provide map data attribution.
L.tileLayer(isRetina ? retinaUrl : baseUrl, {
    attribution: 'Powered by <a href="https://www.geoapify.com/" target="_blank">Geoapify</a> | <a href="https://openmaptiles.org/" rel="nofollow" target="_blank">© OpenMapTiles</a> <a href="https://www.openstreetmap.org/copyright" rel="nofollow" target="_blank">© OpenStreetMap</a> contributors',
    apiKey: myAPIKey,
    maxZoom: 20,
    id: 'osm-bright',
}).addTo(map);

var greenIcon = L.icon({
    iconUrl: "/packages/public/marker.png",

    iconSize:     [25, 25], // size of the icon
    iconAnchor:   [22, 94], // point of the icon which will correspond to marker's location
});

L.marker([51.5, -0.09], {icon: greenIcon}).addTo(map);

var greenIcon = L.icon({
    iconUrl: "/packages/public/marker.png",

    iconSize:     [30, 30], // size of the icon
    iconAnchor:   [22, 94], // point of the icon which will correspond to marker's location
});

L.marker([51.5, -0.09], {icon: greenIcon}).addTo(map);
