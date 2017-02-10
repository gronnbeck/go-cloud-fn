# go-cloud-fn

A Go thing designed to test Google cloud functions https://cloud.google.com/functions/docs/quickstart.

Write your code using go and compile it to server-side node required by Google cloud functions.

Currently, this only implements a http function (which essentially are express handler functions). The idea of this
project is really to show that a pretty simple setup like this (very little dependencies and config required) enables
you to write code for cloud functions using Go <3!

Run `npm run build` to compile the project.

then, you will run:

`gcloud alpha functions deploy helloGO --stage-bucket <your_bucket> --trigger-http`

## License

Copyright © 2017 Martin Sahlen

Distributed under the MIT License