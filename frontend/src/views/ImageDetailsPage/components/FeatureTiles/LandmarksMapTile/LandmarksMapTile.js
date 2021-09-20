import React from 'react';
import { FeatureTemplate, FeatureItem } from 'assets/styles/style';
import { fetchMapAPIConfig } from 'API';




const getMarks = async (landmarks) =>{
  let mapURL = "https://maps.googleapis.com/maps/api/staticmap?&zoom=12&size=250x250&maptype=roadmap&markers=color:red%7Clabel:A%7C"
const mapKey = await fetchMapAPIConfig();
  let x = landmarks?.map((obj) =>
  {return x = obj.latitude}
);
let y = landmarks?.map((obj) =>
{return y = obj.longitude}
);
const finalUrl = mapURL + x[0] +','+y[0] + mapKey;
return await finalUrl;
};

export const LandmarksMapTile = ({ landmarks }) => {

  return (
    <FeatureTemplate>
      <div class="landmark-map">
        <img alt="map" src={getMarks(landmarks)}></img>
      </div>
    </FeatureTemplate>
  );
};
//https://maps.googleapis.com/maps/api/staticmap?&zoom=1&size=400x250&maptype=roadmap&markers=color:red%7Clabel:A%7C40.6892494,-74.04450039999999&key=AIzaSyDiJ71iuXZwLJcP2WcISaRqTZqfjVwsCcQ
//https://maps.googleapis.com/maps/api/staticmap?&zoom=1&size=400x250&maptype=roadmap&markers=color:red%7Clabel:A%7C40.6892494,-74.04450039999999&key=AIzaSyBHNN1GekXMG3npwF6h-sCHGeIJ13E_GSc