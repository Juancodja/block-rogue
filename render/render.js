const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

const player_id = 0; // ID del jugador, se puede cambiar según sea necesario

let enemies = [];
let players = [];
let projectiles = [];

function drawEntities() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  enemies.forEach(e => {
    ctx.fillStyle = e.color;
    ctx.fillRect(e.x, e.y, e.width, e.height);
  });
  players.forEach(p => {
    ctx.fillStyle = p.color;
    ctx.fillRect(p.x, p.y, p.width, p.height);
  });
  projectiles.forEach(p => {
    ctx.fillStyle = p.color;
    ctx.fillRect(p.x, p.y, p.width, p.height);
  });
  
}
console.log(enemies[0]) 

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
      let newEnemies = data.enemies || [];
      if (Array.isArray(newEnemies)) {
        enemies = newEnemies;
      } else {
        console.error("Received data is not an array:", newEnemies);
      }
      let newPlayers = data.players || [];
      if (Array.isArray(newPlayers)) {
        players = newPlayers;
      } else {
        console.error("Received players data is not an array:", newPlayers);
      }
      let newProjectiles = data.projectiles || [];
      if (Array.isArray(newProjectiles)) {
        projectiles = newProjectiles;
      } else {
        console.error("Received proyectiles data is not an array:", newProjectiles);
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
    let player = players.find(p => p.id === player_id);

    socket.send(JSON.stringify({
      player_id: player_id ,
      type: 'attack',
      action: { 
        source_x: player ? player.x : source_x,
        source_y: player ? player.y : source_y,
        target_x: target_x,
        target_y: target_y
      }
    }
    ));
  }
});

const keysPressed = new Set();

document.addEventListener('keydown', (event) => {
  keysPressed.add(event.key.toLowerCase());
});

document.addEventListener('keyup', (event) => {
  keysPressed.delete(event.key.toLowerCase());
});

setInterval(() => {
  if (!socket || socket.readyState !== WebSocket.OPEN) return;

  let dx = 0;
  let dy = 0;

  if (keysPressed.has("arrowup") || keysPressed.has("w")) dy -= 1;
  if (keysPressed.has("arrowdown") || keysPressed.has("s")) dy += 1;
  if (keysPressed.has("arrowleft") || keysPressed.has("a")) dx -= 1;
  if (keysPressed.has("arrowright") || keysPressed.has("d")) dx += 1;

  socket.send(JSON.stringify({
    player_id: 0,
    type: 'move',
    action: { dx, dy }
  }));
}, 1000 / 60); // 30 FPS de envío de movimiento

connectWebSocket();
