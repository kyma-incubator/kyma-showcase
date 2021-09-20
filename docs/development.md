# Development Guide

## Project structure

Kyma Showcase consists of three main components. 

- Frontend with the user interface,
- Backend with the Redis server,
- Serverless Functions.

## Installation

For the installation guidelines, refer to [Installation](../README.md#Installation) section in the main README.md file.

## Configuration

To configure Kyma Showcase, you need to apply the relevant parameters in [secrets.yaml](../resources/secrets.yaml) and fill the config domain in the [frontend-config.yaml](../resources/frontend/frontend-config.yaml) [frontend/apirule.yaml](../resources/frontend/apirule.yaml) and [backend/apirule.yaml](../resources/backend/apirule.yaml) files.


| Parameter | Description |
|-----------|-------------|
| GCP_API_KEY | Your GCP API key. For more information on the key generation refer to the [official documentation](https://cloud.google.com/docs/authentication/api-keys#creating_an_api_key). |
| GCP_EMAIL | Mandatory credential for the GCP access.|
| REDIS_PASSWORD | Required for the authorization layer. |
| API_URL | Defines the API URL that Frontend makes calls to. |
| EVENT_URL | Event publisher URL. |
|API_KEY_PARAM | Your GCP Maps API key. For more information on the key generation refer to the [official documentation](https://console.cloud.google.com/google/maps-apis/start). |
