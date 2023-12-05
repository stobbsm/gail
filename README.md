# Gail: an smtp server written in go

## Why?
1. I want a challenge
2. SMTP either works or it doesn't. There is no in between.
3. Want to modernize the configuration of smtp servers to make them more
   scalable, both in storage and performance.
4. Why not? :D

## Technologies to use
1. Go programming language
2. REST API for external management
3. gRPC for IPC
4. Container first design
5. Scalable database, primarily postgres, with the ability to replace the
   database via interfaces.
   - Should be able to use SQLite for development and small deployments.
6. Security minded: No ability to use as an open relay, must have some form of
   auth supported by the SMTP standards.
7. MUST implement https://www.rfc-editor.org/rfc/rfc5321.html
8. Must NOT run as root unless forced

## Development methodologies
1. Not microservices. Should not need to spin up a kubernetes cluster to use it,
   but you should be able to run it in kubernetes.
2. By nessecity of being SMTP, client-server is a must, where any smtp client
   should be able to communicate with the server
3. Configuration data should be stored in a database, with the ability to
   rollback configuration
4. No stateful configuration files, should make use of environment variables and
   CLI flags for configuration
5. The closest thing to written configuration should be a file containing
   environment variables that are set either via systemd (on linux), or a pre
   command that exports the variables in the scope of the server
