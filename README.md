# Media Server

Accepts http requests for a video identifier, and serves the video, with the
best content type I can guess.

Requests should arrive on `/:id`.
Files should be stored in `/media/:id(.:ext)`.

We use the `file` command to determine the mime type.
