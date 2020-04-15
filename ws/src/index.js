// Import required libraries
const nsq = require('nsqjs');
const http = require('http');
const socketio = require('socket.io');

// Constants
const NSQ_TOPIC = 'AddQuery';

// Create HTTP server
const server = http.createServer();

// Create socket.io wrapper
const io = socketio(server, {
  path: '/',
  serveClient: false,
});

// Create NSQ reader
const reader = new nsq.Reader(NSQ_TOPIC, 'default', {
  lookupdHTTPAddresses: 'nsqlookupd:4161',
});

// Connect to nsqlookupd
reader.connect();

// Process NSQ messages
reader.on('message', message => {
  io.sockets.emit(NSQ_TOPIC, message.body.toString());
  message.finish();
});

// Handle new socket.io connections
io.on('connection', connection => {
  connection.join('default');
});

// Start listening for HTTP requests
server.listen(80, '0.0.0.0', () => {
  console.log('Listening on *:80');
});
