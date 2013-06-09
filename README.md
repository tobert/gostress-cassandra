What is this?
-------------

This is a quick tool I threw together for hammering mercilessly
on Cassandra clusters. It is written in Go and uses the gossie
client library.

But why?
--------

No good reason. I felt like flexing my Go a little and building
something specific to the test I wanted. This is it.

Building
--------

Set up your GOPATH as found in the Go docs. At a minimum:

    export GOPATH=~/.go

Then download gossie:

    go get github.com/carloscm/gossie/src/gossie

Now build:

	git clone git://github.com/tobert/gostress.git
    cd gostress
	go build
	./gostress -list mycluster -mode write
	./gostress -list mycluster -mode read

Setup
-----

Set up the schema with cassandra-cli:

    cassandra-cli --file create_keyspace.cass

Create a cluster list with the ".txt" extension. This can have
just one node of your cluster or all of them. One host:port per line.

    cat > mycluster.txt <<EOF
	cassandra1.foobar.com:9160
	cassandra2.foobar.com:9160
	cassandra3.foobar.com:9160
	cassandra4.foobar.com:9160
	cassandra5.foobar.com:9160
	EOF

Write the data into Cassandra (could take a while):

    ./gostress -list mycluster -mode write

Run the load:

    ./gostress -list mycluster -mode read

