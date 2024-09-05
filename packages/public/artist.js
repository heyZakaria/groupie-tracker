// The Leaflet map Object
var map = L.map('my-map').setView([0,0], 5);

var la 
var lo
// Array of coordinates (lat, lon)
var coordinates = [
    [36.7014631, -118.755997], // Coordinate 1
    [39.5158825, -116.853722], // Coordinate 2
    [32.3293809, -83.1137366]  // Coordinate 3
];


// Create a bounds object from the coordinates
var bounds = L.latLngBounds(coordinates);

// Adjust the map view to fit the bounds of the coordinates
map.fitBounds(bounds);

// Get your own API Key on https://myprojects.geoapify.com
// var myAPIKey = "ab95d9aafa1d4549323898be803875e";

// Retina displays require different map tiles quality
var isRetina = L.Browser.retina;

var baseUrl = "https://maps.geoapify.com/v1/tile/osm-bright/{z}/{x}/{y}.png?apiKey=1ba7c108667243f0ba279f68c82e9b86";
var retinaUrl = "https://maps.geoapify.com/v1/tile/osm-bright/{z}/{x}/{y}@2x.png?apiKey=1ba7c108667243f0ba279f68c82e9b86";

// Add map tiles layer. Set 20 as the maximal zoom and provide map data attribution.
L.tileLayer(isRetina ? retinaUrl : baseUrl, {
    attribution: ' <a href="https://openmaptiles.org/" rel="nofollow" target="_blank"></a> ',
    maxZoom: 9,
    minZoom: 1,
    id: 'osm-bright',
}).addTo(map);

var greenIcon = L.icon({
    iconUrl: "/packages/public/marker.png",
    iconSize: [25, 25], // size of the icon
    iconAnchor: [12.5, 12.5], // center of the icon
});


L.marker([36.7014631, -118.755997], {icon: greenIcon}).addTo(map);


var greenIcon = L.icon({
    iconUrl: "/packages/public/marker.png",
    iconSize: [25, 25], // size of the icon
    iconAnchor: [12.5, 12.5], // center of the icon
});

L.marker([39.5158825, -116.853722], {icon: greenIcon}).addTo(map);


var greenIcon = L.icon({
    iconUrl: "/packages/public/marker.png",
    iconSize: [25, 25], // size of the icon
    iconAnchor: [12.5, 12.5], // center of the icon
});

L.marker([32.3293809, -83.1137366], {icon: greenIcon}).addTo(map);