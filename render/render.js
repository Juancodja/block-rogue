const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

let entities = [];
let players = [];
let proyectiles = [];

function drawEntities() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  entities.forEach(e => {
    ctx.fillStyle = e.color;
    ctx.fillRect(e.x, e.y, e.width, e.height);
  });
  players.forEach(p => {
    ctx.fillStyle = p.color;
    ctx.fillRect(p.x, p.y, p.width, p.height);
  });
  proyectiles.forEach(p => {
    ctx.fillStyle = p.color;
    ctx.fillRect(p.x, p.y, p.width, p.height);
  });
  
}
console.log(entities[0]) 

setInterval(drawEntities, 1000 / 60); // 60 FPS

let socket;
function connectWebSocket() {
  socket = new WebSocket("ws://localhost:8888/ws");

  socket.addEventListener("open", () => {
    console.log("WebSocket connected");
  });

  socket.addEventListener("message", (event) => {
    try {
      const data = JSON.parse(event.data);
      let newEntities = data.entities || [];
      if (Array.isArray(newEntities)) {
        entities = newEntities;
      } else {
        console.error("Received data is not an array:", newEntities);
      }
      let newPlayers = data.players || [];
      if (Array.isArray(newPlayers)) {
        players = newPlayers;
      } else {
        console.error("Received players data is not an array:", newPlayers);
      }
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

canvas.addEventListener('mousedown', (event) => {
  const rect = canvas.getBoundingClientRect();
  const target_x = event.clientX - rect.left;
  const target_y = event.clientY - rect.top;
  const source_x = 500; 
  const source_y = 400; 
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify({
      player_id: 0 ,
      type: 'attack',
      action: { 
        source_x: source_x,
        source_y: source_y,
        target_x: target_x,
        target_y: target_y
      }
    }
    ));
  }
});





connectWebSocket();
