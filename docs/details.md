# Details

Kyma Showcase consists of three main components that work together to provide the final result.

[Frontend](../frontend) presents the user with a clickable drag and drop field and a feed of all previously uploaded pictures. After uploading the desired images, they get added one by one to the feed. Clicking any image takes the user to its details page showing all the information acquired from processing the image.

[Backend](../backend) manages all calls to the Redis Database. Upon successfully saving an image, Backend sends an initial event triggering serverless Functions.

[Functions](../resources/functions) work with the image ID received from the event to acquire the right image from the Redis Database and send it to the Google Cloud Platform. The Functions update the Redis Database entry with newly acquired information and depending on its content they can send events triggering more Functions in order to get further details on the image.
