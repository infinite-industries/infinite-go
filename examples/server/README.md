# server
A rudimentary HTTP server to represent II events / venues in JSON-LD form.

Starting:
```console
$ ./server --loglevel debug
{"level":"info","service":"ii","time":"2023-07-25T11:06:45-04:00","message":"Log level set to debug"}
{"level":"debug","service":"ii","time":"2023-07-25T11:06:45-04:00","message":"config initialized: address=:7070, prefix_path=/, loglevel=debug"}
{"level":"debug","service":"ii","time":"2023-07-25T11:06:45-04:00","message":"binding middleware.Hearbeat to /ping"}
{"level":"debug","component":"jsonld","message":"service bound to /"}
{"level":"info","service":"ii","time":"2023-07-25T11:06:45-04:00","message":"starting server on :7070"}
```

Example Requests
* http://localhost:7070/event/4da9761e-1dfd-488c-8f53-2a95cac188c1
* http://localhost:7070/venue/89cd986b-382b-4bf3-95b4-9b2e42db335b
