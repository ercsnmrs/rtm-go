# RTM-GO

Sample implementation of Real Time Monitoring Using Websockets

## Build

Build in the normal way on Mac/Linux:

~~~
go build -o rtm-go cmd/web/*.go
~~~

Or on Windows:

~~~
go build -o rtm-go.exe cmd/web/.
~~~

Or for a particular platform:

~~~
env GOOS=linux GOARCH=amd64 go build -o rtm-go cmd/web/*.go
~~~

## Requirements

rtm-go requires:
- Postgres 11 or later (db is set up as a repository, so other databases are possible)
- An account with [Pusher](https://pusher.com/), or a Pusher alternative 
(like [ipê](https://github.com/dimiro1/ipe))

## Run

First, make sure ipê is running (if you're using ipê):

On Mac/Linux
~~~
cd ipe
./ipe 
~~~

On Windows
~~~
cd ipe
ipe.exe
~~~

Run with flags:

~~~
./rtm-go \
-dbuser='tcs' \
-pusherHost='localhost' \
-pusherPort='4001' \
-pusherKey='123abc' \
-pusherSecret='abc123' \
-pusherApp="1" \
-pusherSecure=false
~~~~

## All Flags

~~~~
./rtm-go -help
Usage of ./rtm-go:
  -db string
        database name (default "rtmsvc")
  -dbhost string
        database host (default "localhost")
  -dbport string
        database port (default "5432")
  -dbssl string
        database ssl setting (default "disable")
  -dbuser string
        database user
  -dbpass string
        database user
  -domain string
        domain name (e.g. example.com) (default "localhost")
  -identifier string
        unique identifier (default "rtm-go")
  -port string
        port to listen on (default ":4000")
  -production
        application is in production
  -pusherApp string
        pusher app id (default "9")
  -pusherHost string
        pusher host
  -pusherKey string
        pusher key
  -pusherPort string
        pusher port (default "443")
  -pusherSecret string
        pusher secret
   -pusherSecure
        pusher server uses SSL (true or false)

Sample ./rtm-go -dbhost localhost -dbport 5432 -dbpass pass -dbuser user

~~~~