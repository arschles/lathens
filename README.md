# lathens

[WIP] A caching proxy for the `/@v/list` and `/@latest` endpoints for Athens.

This is useful in situations where users of a proxy might run a lot of `go get` commands without versions on them. Similarly, a user might add a new `import` in their code and run a `go build`. In these cases (and others!), these endpoints will get hit.

[Athens](https://github.com/gomods/athens) gets up-to-date information from the upstream VCS when handling both the `/@v/list` and `/@latest` endpoints. If you cache the values for these endpoints, you can reduce Athens' dependence on the upstream VCS systems. If you store the values forever, you can completely isolatet Athens from the upstream.

In most situations I've seen so far, caching or storing `/@v/list` and `/@latest` isn't really needed. There are two where the isolation is really handy:

- You're running a public proxy like [athens.azurefd.net](https://athens.azurefd.net) and you don't know when and how many "bare" `go get` commands will hit the server
  - Caching these values helps prevent Athens from making tons of VCS requests
- You're running an Athens deployment that is completely shut off from the public internet, and want to pre-seed the `/@v/list` and `/@latest` values


>You can use this with any other Go module proxy server as well
