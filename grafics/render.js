import { drawEnemyAnimated } from './render/enemyAnimation.js';
import { Entity } from './entity.js'


const VIEW_WIDTH = 1280
const VIEW_HEIGHT = 800


const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');
const pre = document.getElementById('infoPre')


const button = document.getElementById("changeWeapon")


const player_id = 0; 

let enemies = new Map();
let players = new Map();
let projectiles = new Map();


const playerSprite = new Image();
playerSprite.src = 'sprites/player.png'

const background = new Image();
background.src = 'sprites/background.png'




let socket;
function connectWebSocket() {
    socket = new WebSocket("ws://localhost:8888/ws");

    socket.addEventListener("open", () => {
        console.log("WebSocket connected");
    });

    socket.addEventListener("message", (event) => {
        try {
            const data = JSON.parse(event.data);

            const newEnemies = new Map(Object.entries(data.enemies || {}));
        
            enemies.forEach((entity, key) => {
                entity.dead = true
            })
            newEnemies.forEach((entity,key) =>{
                if(enemies.has(key)){
                    let enemy = enemies.get(key) 
                    enemy.updateFromServer(entity)
                }
                else{
                    enemies.set(key, new Entity(entity))
                }
            })
            console.log(enemies)

            const toDelete = []
            enemies.forEach((entity, key) => {
                if(entity.dead === true){
                    toDelete.push(key)
                }
            })
            
            for(const uuid of toDelete){
                enemies.delete(uuid)
            }

            console.log(enemies.size)
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
        let player_id = 0
        socket.send(JSON.stringify({
            player_id: player_id,
            type: 'attack',
            action: {
                player_id: player_id,
                target_x: target_x,
                target_y: target_y,
                source_x: player ? player.x : source_x,
                source_y: player ? player.y : source_y,
                width: 40,
                height: 40, 
                damage: 40, 
                distance_from_source: 10,
                time_alive: 10,
                speed: 0, 
                max_distance: 10000
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

function makeInfo(){
    let info = {}
    return info
}


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

setInterval(()=>{
    pre.textContent = JSON.stringify(makeInfo(), null, 2)
}, 1000)


function renderLoop(now) {
  for (const [, entity] of enemies) {
    entity.updateAnimation(now);
  }

  drawAllEntities(ctx);
  requestAnimationFrame(renderLoop);
}

function drawAllEntities(ctx) {
  ctx.clearRect(0, 0, canvas.width, canvas.height);

  for (const [, entity] of enemies) {
    entity.draw(ctx);
  }

//   for (const [, projectile] of projectiles) {
//     projectile.draw(ctx);
//   }

//   for (const [, player] of players) {
//     player.draw(ctx);
//   }
}



requestAnimationFrame(renderLoop);

connectWebSocket();
