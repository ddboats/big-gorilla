# big-gorilla
Provides a simple interface and middle-ware to analyze DNS requests.

# Getting Started

### Installation
This project only requires docker and docker-compose. To starting using it, simply download this repository to your machine.

Example:
```bash
git clone https://github.com/ddboats/big-gorilla
cd ./big-gorilla/
```

### Configuring
Inside the `docker-compose.yml` file you can find the service `http`. You may change the publish port value to whatever value you prefer to access the web interface from.

Also within `docker-compose.yml` you can find a service `dnsd`. The publish value may also be changed to whatever port you'd like to access the DNS server middleware from. Due note that to utilize it on a standard network configuration, port `53` is to be desired.

### Running
To start the big-gorilla application, all you need to do is type `docker-compose up` and compose will build and start each required service.

Once it is running you can access the web interface by going to `http://SERVER_ADDRESS_HERE/`.

You may notice that there is nothing on the page. This is expected. As of now, big-gorilla only adds data once the web app begins listening to it. In the future, persistence may be supported.

Whenever a request is sent to the DNS middleware server, you will see a new entry added to the web interface.
