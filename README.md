<h1 align="center">SunSend s simple and fast message board</h1>

<h3>Notice: it's still in progress</h3>

<hr>

### How it works ?

Welcome to the world of Message Boards :-)

There is 2 program for using sunsend:

1. Server Side Program (backend): A program that can provide the chat system written in go (in progress)
2. Client Side program (frontend): A program that can provide a interface for using server side program written in JS (soon)

For sending messages, you need to specify the special channel, each channel has some ID that you can find in the API documentation (in progress)

This repo is a server side Rest API that provide the chat system, if you wanna just use the chat in your website, you need to see the client repo (in progress)

For more things, you can checkout TODO.md
<br/>
<br/>
<br/>

### How does API KEY works?

SunSend server has Client (or you can write one) and for connecting to the server, you need to select your **special** server

And you wanna sure that just **your** client can send messages to the your own **server**, API KEY can provide a authentication process that

You can use on your client and just you (or your website AKA your users) can send message to your SunSend server.

You can generate a api key with cmd/gen*key.go and put that to your .env file as \_API_KEY* variable

for generating api key:

```makefile
make setup
```

<br/>
<br/>

### How to run the server?

For running SunSend you need to do somethings before running it
<br/>

_First_ of All, you need to generate a api key for your server, you can read more about it on the _How Does API Works_ part

_Second_, you need to fill your .env file with your own values, notice that that file is a very important file and you _must_
fill it with your own values, variables are:

```bash
PORT: your default port for your server
KEY: your API KEY for your sevrer
```

_Thirth_, you need to config your server, if you wanna use the default configuration type:

```makefile
make config
```

this will install a simple configuration for you

_Fourth_, you can finnally run your server with command:

```makefile
make run
```

<br/>
<br/>
<br/>

**Build and run SunSend (soon)**
