import React from 'react';
import { FeatureTemplate } from 'assets/styles/style';
import { fetchMapAPIConfig } from 'API';

const getMarks = async (landmarks) => {
  let mapURL = 'https://maps.googleapis.com/maps/api/staticmap?&zoom=12&size=250x250&maptype=roadmap&markers=color:red%7Clabel:A%7C';
  const mapKey = await fetchMapAPIConfig();
  let x = landmarks?.map((obj) => {
    return (x = obj.latitude);
  });
  let y = landmarks?.map((obj) => {
    return (y = obj.longitude);
  });
  const finalUrl = mapURL + x[0] + ',' + y[0] + mapKey;
  console.log(finalUrl)
  return finalUrl;
};

export const LandmarksMapTile = ({ landmarks }) => {
  const map = getMarks(landmarks);
  console.log(map);
  return (
    <FeatureTemplate>
      <div class="landmark-map">
        {map && <img alt="map" src={map}></img>}
      </div>
    </FeatureTemplate>
  );
};
