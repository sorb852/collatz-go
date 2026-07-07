# Collatz conjecture API

> Oh hey look its one of my favourite things.

This API is used to generate the sequence of numbers using the Collatz rules given starting point

## Running

Firstly, clone the repository.

```sh
git clone https://gitub.com/sorb852/collatz-go
cd collatz-go
```

Then we can either build

```sh
go build
./collatz-go
```

or just run freely

```sh
go run .
```

## Usage

### Requesting

The API has one endpoint (lol) being `GET /collatz?n=[NUM]`.

For instance, running this using curl would look like this:

```sh
curl localhost:8080/collatz?n=5
```

### Response

The response is sent as `JSON` and either comes through as `CollatzResponse` or `CollatzResponseError`.

In the case of an error, the body will look like this:

```json
{
    "message": "some kind of error happened"
}
```

(**Note** that the sent back json is not beatified. It comes back in a clumped up form.)

And when it succeeds without issues, the body will look like this:

```json
{
    "starting_nubmer": 852,
    "length": 25,
    "peak": 2112,
    "chain": [852, 426, ...]
}
```

(**Note** that this is just the form of the response, this is **NOT** true)

Maybe in the future I will add more stats. Though it doesnt happen that often so don't get your hopes up

## So, why?

I merely needed an excuse to learn go, before this I didn't like the syntax and gave up. Though I gave it a second chance because it always looked cool to me.

Choosing the *Collatz Conjecture* is just the thing I do when I learn any new language. And this time I put it up with Go's exceptional networking library (yes it is that good).

<!-- vim: wrap -->
