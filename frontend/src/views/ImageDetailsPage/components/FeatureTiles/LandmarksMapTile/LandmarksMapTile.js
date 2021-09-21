import React, { useEffect, useState } from 'react';
import { FeatureTemplate } from 'assets/styles/style';
import { fetchMapAPIConfig } from 'API';

export const LandmarksMapTile = ({ landmarks }) => {
  const [mapUrl, setMapUrl] = useState('');

  useEffect(() => {
    const getMarks = async (landmarks) => {
      let mapURL = 'https://maps.googleapis.com/maps/api/staticmap?&zoom=12&size=490x260&maptype=roadmap&markers=color:red%7Clabel:A%7C';
      const mapKey = '&key=' + (await fetchMapAPIConfig());
      let x = landmarks[0].latitude;
      let y = landmarks[0].longitude;
      const finalUrl = mapURL + x + ',' + y + mapKey;
      setMapUrl(finalUrl);
    };
    getMarks(landmarks);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <FeatureTemplate>{mapUrl && <img alt="map" src={mapUrl} />}</FeatureTemplate>;
};
