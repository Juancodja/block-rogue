const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

let entities = [
  { x: 50, y: 60, width: 100, height: 100, color: 'red' },
  { x: 200, y: 150, width: 80, height: 120, color: 'blue' },
  { x: 400, y: 300, width: 150, height: 60, color: 'green' },
];

function drawEntities() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  entities.forEach(e => {
    ctx.fillStyle = e.color;
    ctx.fillRect(e.x, e.y, e.width, e.height);
  });
}

setInterval(drawEntities, 100);

let socket;
function connectWebSocket() {
  socket = new WebSocket("ws://localhost:8888/ws");

  socket.addEventListener("open", () => {
    console.log("WebSocket connected");
  });

  socket.addEventListener("message", (event) => {
    try {
      console.log("Message received:", event.data);
      const data = JSON.parse(event.data);
      entities = [data]; // remplaza entidades con la nueva
    } catch (e) {
      console.error("Error parsing message:", e);
    }
  });

  socket.addEventListener("close", () => {
    console.log("WebSocket closed. Reconnecting in 1 second...");
    setTimeout(connectWebSocket, 1000);
  });

  socket.addEventListener("error", (err) => {
    console.error("WebSocket error:", err);
    socket.close(); // Fuerza cierre para activar reconexión
  });
}

connectWebSocket();
