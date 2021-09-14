# Details

Project consists of three main components that work together to provide the final result.

[Frontend](./frontend) presents the user with a clickable drag and drop field and a feed of all previously uploaded pictures. After uploading desired image it gets added to the feed in order of upload time. Clicking any image takes the user to a details page where they can read all details info acquired.

[Backend](./backend) manages all calls to the database. Upon sucessfully saving the image backend sends an initial event triggering the serverless functions.

[Functions](./resources/functions) work with the image ID received from the event to acquire the right image from the database and then send it to Google Cloud Platform. Functions update the database entry with newly acquired information and depending on its content can send events triggering more functions in order to get further details from the image.
