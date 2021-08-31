#API Contract
****
###Getting all images
Homepage startup triggers downloading all images from database.\
Frontend makes following request\
`GET v1/images`\
The endpoint returns the following JSON \
`[{ id: "#", title: "#", content: base64, status: "#" },{},...]`\
The endpoint returns the following HTTP codes: 
- 200 (if request succeeded)
- 500 (if failed)

###Adding an image
User adds a new image\
`POST v1/images`\
The endpoint returns the following JSON \
`{id: UUID}`\
The endpoint returns the following HTTP codes: 
- 202 (if request accepted)
- 500 (if failed)
- 400 (if bad input)

###Getting an image
User clicks on an image\
`GET v1/images/{id}`\
The endpoint returns the following JSON \
`[{ id: "#",
title: "#",
content: base64,
description: {JSON from GCP},
status: "#"
}]`\
The endpoint returns the following HTTP codes:
- 200 (if processing is done)
- 202 (if processing is in progress)
- 500 (if failed)

###Updating an image
Image details in database are being updated with values (e.g. recognized labels on photo) returned from GCP\
`PUT v1/images/{id}`\
The endpoint returns the following HTTP codes:
- 200 (if processing is done)
- 404 (if record does not exist)
- 500 (if failed)