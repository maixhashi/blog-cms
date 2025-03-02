import React from 'react';
import RouteMapVisualizer from '../components/RouteMapVisualizer';

const RouteMapPage: React.FC = () => {
  return (
    <div className="route-map-page">
      <h1>Blog CMS 画面遷移図</h1>
      <RouteMapVisualizer />
    </div>
  );
};

export default RouteMapPage;
