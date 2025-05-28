export class Entity {
  constructor(entity) {
    this.uuid = entity.uuid;
    this.id = entity.id;
    this.name = entity.name;
    this.x = entity.x;
    this.y = entity.y;
    this.dx = entity.dx;
    this.dy = entity.dy;
    this.color = entity.color;
    this.width = entity.width;
    this.height = entity.height;
    this.health = entity.health;
    this.speed = entity.speed;
    this.traveled_distance = entity.traveled_distance;
    this.max_distance = entity.max_distance;
    this.type = entity.type;
    this.time_alive = entity.time_alive;
    this.max_time_alive = entity.max_time_alive;

    this.frame = 0;
    this.lastUpdate = performance.now();
    this.frameWidth = 64;
    this.frameHeight = 64;
    this.totalFrames = 16;
    this.dead = false; 

    this.sprite = new Image();
    this.sprite.src = `sprites/${this.name}.png`;
    this.sprite.onerror = () => {
      this.sprite.onerror = null;
      this.sprite.src = 'sprites/default.png';
    };
  }

  updateFromServer(entity) {
    this.x = entity.x;
    this.y = entity.y;
    this.dx = entity.dx;
    this.dy = entity.dy;
    this.width = entity.width;
    this.height = entity.height;
    this.color = entity.color;
    this.health = entity.health;
    this.speed = entity.speed;
    this.traveled_distance = entity.traveled_distance;
    this.max_distance = entity.max_distance;
    this.type = entity.type;
    this.time_alive = entity.time_alive;
    this.max_time_alive = entity.max_time_alive;

    this.dead = false
  }

  updateAnimation(now) {
    const dt = now - this.lastUpdate;
    if (dt > 2000) {
      this.frame = (this.frame + 1) % this.totalFrames;
      this.lastUpdate = now;
    }
  }

  draw(ctx) {
    if (!this.sprite || !this.sprite.complete) return;
    const sx = this.frame * this.frameWidth;
    ctx.drawImage(
      this.sprite,
      sx, 0, this.frameWidth, this.frameHeight,
      Math.floor(this.x - this.width / 2),
      Math.floor(this.y - this.height / 2),
      this.width, this.height
    );
  }
}
