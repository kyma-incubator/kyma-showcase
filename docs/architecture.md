# Architecture

The diagram presents the basic workflow of Kyma Showcase.

![Diagram Kyma Showcase](./assets/diagram_showcase.svg)

1. Frontend gets all images stored in the Redis Database and displays them in the feed.
2. A user uploads an image using Frontend's UI.
3. The newly uploaded image is stored in the Redis Database and then displayed in the feed.
4. Backend sends an event containing the Redis Database ID used to access the image.
5. Serverless Functions are triggered by the event and get the image from the Redis Database using the given ID.
6. The Functions send the image in base64 to the Google Cloud Platform (GCP) for processing.
7. The Redis Database entry is updated with newly obtained image details.
8. Depending on the details content, a new event may be sent triggering next Functions.
9. Upon clicking a single image in the Frontend UI, a details page is displayed listing all the information received from the GCP.

The initial event triggers `Multiple object detection`, `Landmark detection` and `General labels` Functions. The remaining ones get triggered based on the content of the response recieved from the GCP.

| Content required | Information type | Serverless Functions triggered |
|-----------|-------------|-------------|
| Text/Font | Label | Text detection |
| Handwriting | Label | Handwriting detection |
| Logo | Label | Logo detection|
| Text | String | Text category detection, Sentiment detection, Text entities and Entity sentiment |
| Person | Object | Face detection |
