# Kyma Showcase
[![Go Report Card](https://goreportcard.com/badge/github.com/kyma-incubator/Kyma-Showcase)](https://goreportcard.com/report/github.com/kyma-incubator/Kyma-Showcase)
[![GoDoc](https://godoc.org/github.com/kgithub.com/kyma-incubator/Kyma-Showcase?status.svg)](https://godoc.org/github.com/kyma-incubator/Kyma-Showcase)

## Overview

Kyma Showcase about object-recognition relying on GCP API and Lambda Workflow. Customer uploads an image on the front page which is then displayed in a feed. Clicking on any image in feed redirects you to details page which displays all the information acquired from processing that image.

## Prerequisites

- [Kyma](https://kyma-project.io/) cluster 
>**TODO:Add version after first release.**
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) to deploy the application

## Installation

1. To run Kyma-Showcase application, you have to clone repository:
```
git clone https://github.com/kyma-incubator/Kyma-Showcase.git
```
2. Get cluster domain and replace placeholder values in following files:
- resources/backend/apirule.yaml
- resources/frontend/apirule.yaml
- resources/frontend/frontend-config.yaml

3. Set secrets by replacing placeholders in file:
- resources/secrets.yaml

4. To deploy application on k8s cluster, run the following commands:
```
kubectl apply -f resources -R
```

5. To get the URL of the application run the following command:
```
kubectl get apirule frontend -o jsonpath='{.spec.service.host}'
```

## Usage

![Diagram Kyma Showcase](./docs/assets/diagram_showcase.svg)

[Frontend](./frontend) presents the user with a clickable drag and drop field and a feed of all previously uploaded pictures. After uploading desired image it gets added to feed in order of upload time. Clicking any image takes the user to a details page where they can read all details info acquired.

[Backend](./backend) receives the image and saves it in a database. Upon sucessfully saving the image event is sent from it triggering the serverless functions.

[Functions](./resources/functions) work with the image ID received from the event to acquire the right image from the database which is then sent to Google Cloud Platform in base64. Function updates the database entry with newly acquired information and depending on its content can send another event triggering more functions in order to get even more details from the image.

## Development

Detailed development guide can be found [here](./docs/01-development.md).

To be able to contribute follow the workflow described in [`git-workflow.md`](https://github.com/kyma-project/community/blob/master/contributing/03-git-workflow.md) document.
