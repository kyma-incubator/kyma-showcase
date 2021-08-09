# Kyma Showcase

## Overview

Kyma Showcase about object-recognition relying on GCP API and Lambda Workflow.

## Prerequisites

TODO: List the requirements to run the project.

## Installation

To run Kyma-Showcase application, run the following commands:
```
git clone https://github.com/kyma-incubator/Kyma-Showcase.git
```

Run Docker app.

Run the project in any IDE.
```
cd backend
make build
```

Go to the .env file and paste the code below:
```
REDIS_URL=redis:6379
REDIS_PASSWORD=“”
PORT=8081
```

In terminal run the folowing commands:

```
docker build -t image .
```
```
docker-compose up --build
```
Now backend is ready, then we have to start frontend application.

```
cd ../frontend
npm ci
npm start
```
## Usage

TODO: Explain how to use the project. You can create multiple subsections. Include the instructions or provide links to the related documentation.

## Development

TODO: Add instructions on how to develop the project. It must be clear what to do and how to trigger the tests so that other contributors know how to make their pull requests acceptable. Include the instructions or provide links to related documentation.
