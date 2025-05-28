const enemySprite = new Image();
enemySprite.src = './sprites/enemy.png';

const FRAME_WIDTH = 64;
const FRAME_HEIGHT = 64;
const TOTAL_FRAMES = 16;

let currentFrame = 0;

export function drawEnemyAnimated(ctx, e, deltaTime) {
  e.animTimer += deltaTime;
  if (e.animTimer >= 100) {
    e.animTimer = 0;
    e.animFrame = (e.animFrame + 1) % TOTAL_FRAMES;
  }

  const sx = currentFrame * FRAME_WIDTH;
  ctx.drawImage(
    enemySprite,
    sx, 0, FRAME_WIDTH, FRAME_HEIGHT,
    Math.floor(e.x - FRAME_WIDTH / 2),
    Math.floor(e.y - FRAME_HEIGHT / 2),
    FRAME_WIDTH, FRAME_HEIGHT
  );
}
