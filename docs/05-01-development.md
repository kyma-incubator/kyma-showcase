---
title: Development guide
type: Configuration
---

# Development guide

## Project structure

Project is divided into three main components. Frontend with user interface, backend with Redis server and serverless functionalities.

## Prerequisites

- [Kyma](https://kyma-project.io/) cluster
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) to deploy the application

## Installation

Installation can be found [here](../README.md#Installation).

## Configuration

To configure Kyma Showcase you need to configure the following parameters in [secrets.yaml](../resources/secrets.yaml) and [frontend-config.yaml](../resources/frontend/frontend-config.yaml).


| Parameter | Description |
|-----------|-------------|
| GCP_API_KEY | Your Google Cloud Platform API key, for more information on key generation refer [here](https://cloud.google.com/docs/authentication/api-keys#creating_an_api_key). |
| GCP_EMAIL | Mandatory credential for Google cloud platform access.|
| REDIS_PASSWORD | Required as a set up for the authorization layer |
| API_URL | Defines the API URL that frontend will make calls to. |


## Usage


TODO