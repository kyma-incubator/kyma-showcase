import React, { useEffect, useState } from 'react';
import { FeatureTemplate } from 'assets/styles/style';
import { fetchMapAPIConfig } from 'API';

export const LandmarksMapTile = ({ landmarks }) => {
  const [mapUrl, setMapUrl] = useState('');

  useEffect(() => {
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
      setMapUrl(finalUrl);
    };
    getMarks(landmarks);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <FeatureTemplate>
        {mapUrl && <img alt="map" src={mapUrl}/>}
    </FeatureTemplate>
  );
};
