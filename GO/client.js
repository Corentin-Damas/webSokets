// Establish the connection with the websocket at an url end point
let socket = new WebSocket("ws://localhost:3000/ws");
let socketOrder = new WebSocket("ws://localhost:3000/order");
// Server cmd message : new incoming connection from client: chrome://new-tab-page

// Do somehting when the serveur send back some message
socket.onmessage = (event) => {
  console.log("received msg from the server:", event.data);
};
socketOrder.onmessage = (event) => {
  console.log("received order from the server:", event.data);
};

// Send a message to the serveur
socket.send("Hello from client");

// Msg from serveur side will display : Hello from client
// Msg from Client side will display : received from the server: thank you for the message!
