import React from 'react';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';

let marker = { lat: 51.505, lng: -0.09 };
let latStr;
let lngStr;
let mapURL = "https://maps.googleapis.com/maps/api/staticmap?&zoom=13&size=900x600&maptype=roadmap&markers=color:blue%7Clabel:S%7C"
const mapKey = "&key=AIzaSyBHNN1GekXMG3npwF6h-sCHGeIJ13E_GSc"

export const LandmarksMapTile = ({ landmarks }) => {
  return (
    <FeatureTemplate>
      {landmarks?.map((obj, i) => (
        <FeatureItem key={i}>
          {(marker.lat = obj.latitude)}
          {(marker.lng = obj.longitude)}
        </FeatureItem>
      ))}
      {latStr = marker.lat.toString()}
      {lngStr = marker.lng.toString()}
      {console.log(latStr)}
      {console.log(lngStr)}
      {mapURL = mapURL + latStr + "," + lngStr + mapKey}
      <div class="landmark-map">
        <img width="600px" height="900px" alt="map" src={mapURL}></img>
      </div>
    </FeatureTemplate>
  );
};