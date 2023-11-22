# Simulation of http connection lost

## server

/hello handle 2 channel
1st is context timeout 3s
2nd is request context

if request context is done means client connection is lost

## client

client play 3 role

1. just call by comment timeout: this scenario should success
2. just call and interrupt before reach 3s: this scenario should see "connection lost" message in server side
3. call with timeout (uncomment): this scenario should see "connection lost" message in server side
