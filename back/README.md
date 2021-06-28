# Cryptio Front Exam / Back

Here is how you can run this back-end:

## Without Docker

If you can/want to run Go code in your machine run:

```
go mod run .
```

And visit
[http://127.0.0.1:8080/swagger/index.html](http://127.0.01:8080/swagger/index.html)
in order to access the specifications of the API.

## With Docker

If you are not comfortable with running Go code, here is how to run this
back-end in Docker...

Build the image:

```
docker build . -t cryptio-front-exam-back
```

Run a container:

```
docker run \
  -p 127.0.0.1:8080:8080 \
  -e GIN_MODE=release \
  -e ADDR=0.0.0.0 \
  -e PORT=8080 \
  cryptio-front-exam-back
```

Finally, visit
[http://127.0.0.1:8080/swagger/index.html](http://127.0.01:8080/swagger/index.html)
in order to access the specifications of the API.
