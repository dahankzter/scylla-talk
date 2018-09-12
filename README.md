# Scylla Talk - Some hands on stuff...

This project has a few hands on exercises to learn how to use Scylla as a programmer.
You can pick between writing your code in Go, Java or Python.
Should you by any chance want to  use another language you are welcome to do so but any
help you might need is then restricted to Scylla and CQL.

## Mission

Your mission, should you choose to accept it, is to modify and extend one or several of the
provided programs with the following.

1. Have the program accept an argument in the form of a UUID and print just this row.
2. Have the program accept an argument in the form of a UUID, an integer denoting the column and 
a value to update the object that this UUID corresponds to and with this new value.
3. Add another table called `comments` with three columns: id, ts and comment.
This table should store comments made by certain user (id) at a given time (ts).
4. Create crud operation to insert a comment for a user.
5. Implement a way to list the comments for just one given user.

## Setup Scylla

1. Install [Scylla Tools](https://github.com/scylladb/scylla-tools-java) and put the `scylla-tools-java/bin` in your $PATH
2. Install [Docker](https://www.docker.com/) using your package manager of choice
3. Execute `docker run --name some-scylla -d scylladb/scylla` to start Scylla
4. Find the IP (<HOST> from now on) number of the container like this: `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' some-scylla`
5. Initialize the database:
```bash
cqlsh -f schema/001-init_schema.cql <HOST>
cqlsh -f ../schema/002-add_some_users.cql <HOST>
```

## Setup for Go

1. Install [Go](https://golang.org/) using your package manager of choice
2. Go to the directory ___scylla-talk-go___
3. Execute `go build` to produce the program executable
4. Run the program like this:

```bash
./scylla-talk-go -host <HOST>
```

## Setup for Java

1. Install [Maven](https://maven.apache.org/) using your package manager of choice
2. Go to the directory ___scylla-talk-java___
3. Execute `mvn package` to produce the program executable
4. Run the program like this:

```bash
java -jar target/scylla-talk-1.0-SNAPSHOT.jar <HOST>
```

## Setup for Python

1. Install [Python](https://www.python.org/) using your package manager of choice
2. Go to the directory ___scylla-talk-python___
3. Install the cassandra driver `sudo pip install cassandra-driver` or without sudo `pip install --user cassandra-driver`
4. Run the program like this:

```bash
python cluster.py <HOST>
```
