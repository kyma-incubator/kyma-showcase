# Kyma Showcase
[![Go Report Card](https://goreportcard.com/badge/github.com/kyma-incubator/Kyma-Showcase)](https://goreportcard.com/report/github.com/kyma-incubator/Kyma-Showcase)
[![GoDoc](https://godoc.org/github.com/kgithub.com/kyma-incubator/Kyma-Showcase?status.svg)](https://godoc.org/github.com/kyma-incubator/Kyma-Showcase)

## Overview

Kyma Showcase is a project about object-recognition relying on GCP API and Serverless Functions Workflow. A customer uploads an image on the front page which is then displayed in a feed. Clicking on any image in feed redirects you to its details page which displays all the information acquired from processing that image.

## Prerequisites

- [Kyma](https://kyma-project.io/) cluster 
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) to deploy the application

## Installation

1. To run Kyma Showcase application, clone the repository:
```
git clone https://github.com/kyma-incubator/Kyma-Showcase.git
```
2. Get cluster domain and replace placeholder values in the following files:
- `resources/backend/apirule.yaml`
- `resources/frontend/apirule.yaml`
- `resources/frontend/frontend-config.yaml`

3. Set secrets by replacing placeholders in the file:
- `resources/secrets.yaml`

4. To deploy application on Kubernetes cluster, run the following command:
```
kubectl apply -f resources -R
```

5. To get the application URL run the following command:
```
kubectl get apirule frontend -o jsonpath='{.spec.service.host}'
```

## Development

Detailed development guide can be found [here](./docs/development.md).

To be able to contribute follow the workflow described in [`git-workflow.md`](https://github.com/kyma-project/community/blob/master/contributing/03-git-workflow.md) document.
