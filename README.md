# Kyma Showcase

## Overview

Kyma Showcase about object-recognition relying on GCP API and Lambda Workflow.

## Prerequisites

- Kubernetes cluster
- Kyma installed on cluster

## Installation

1. To run Kyma-Showcase application, you have to clone repository:
```
git clone https://github.com/kyma-incubator/Kyma-Showcase.git
```
2. Get cluster domain and replace placeholder values in following files:
- resources/backend/apirule.yaml
- resources/frontend/apirule.yaml
- resources/frontend/frontend-config.yaml

3. Set secrets with command:
```
kubectl create secret generic kyma-showcase-secret --from-literal=REDIS_PASSWORD={YOUR_PASSWORD}
```

4. To deploy application on k8s cluster, run the following commands:
```
kubectl apply -f resources -R
```

5. To get the URL of the application run the following command:
```
kubectl get apirule frontend -o jsonpath='{.spec.service.host}'
```

## Usage

TODO: Explain how to use the project. You can create multiple subsections. Include the instructions or provide links to the related documentation.

## Development

TODO: Add instructions on how to develop the project. It must be clear what to do and how to trigger the tests so that other contributors know how to make their pull requests acceptable. Include the instructions or provide links to related documentation.
