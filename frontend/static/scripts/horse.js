const HORSE_DIRECTIONS = {
  RIGHT: 1,
  LEFT: 2,
  UP: 3,
  DOWN: 4
}

class Horse {
  
  
  /**
   * 
   * @param {x: number, y: numer, sprite: string} setup
   */
  constructor({width, height, sprite, screenWidth, screenHeight, x = 0, y = 0, direction = HORSE_DIRECTIONS.RIGHT}) {
    this._x = x;
    this._y = y;
    this._width = width;
    this._height = height;
    this._screenWidth = screenWidth - width;
    this._screenHeight = screenHeight - height;
    this._id = `horse-${Math.random() * 10}`;
    this._speed = 10;
    this._sprite = sprite;
    if (x == 0 && y == 0){
      direction = HORSE_DIRECTIONS.DOWN;
    }
    this._direction = direction;
    this._buildElement();
  }

  htmlElement() {
    return this._element;
  }

  _buildElement() {
    this._element = document.createElement("img");
    this._element.id = this._id;
    this._element.src = this._sprite;
    this._element.classList.add("running-horse");
    this._updateStyles();
    return this._element;
  }

  update() {
    this.move();
  }

  move() {
    this._calculateMovement();
    this._updateDirection();
    this._updateStyles();
  }

  _calculateMovement() {
    switch(this._direction) {
      case HORSE_DIRECTIONS.RIGHT:
        this._x += this._speed;
        break;
      case HORSE_DIRECTIONS.LEFT:
        this._x -= this._speed;
        break;
      case HORSE_DIRECTIONS.UP:
        this._y -= this._speed;
        break;
      case HORSE_DIRECTIONS.DOWN:
        this._y += this._speed;
        break;
    }
  }

  _updateDirection() {
    if (this._x > this._screenWidth) {
      this._x = this._screenWidth;
      this._direction = HORSE_DIRECTIONS.UP;
      return;
    } 

    if (this._x > 0 && this._y <= 0) {
      this._y = 0;
      this._direction = HORSE_DIRECTIONS.LEFT;
      return;
    }

    if (this._y == 0 && this._x <= 0) {
      this._x = 0;
      this._direction = HORSE_DIRECTIONS.DOWN;
      return;
    }

    if (this._x == 0 && this._y >= this._screenHeight) {
      this._y = this._screenHeight;
      this._direction = HORSE_DIRECTIONS.RIGHT;
    }
  }

  _rotateSprite() {
    switch(this._direction) {
      case HORSE_DIRECTIONS.LEFT:
        this._element.style.rotate = "180deg";
        break;
      case HORSE_DIRECTIONS.RIGHT:
        this._element.style.rotate = "0deg";
        break;
      case HORSE_DIRECTIONS.UP:
        this._element.style.rotate = "-90deg";
        break;
      case HORSE_DIRECTIONS.DOWN:
        this._element.style.rotate = "90deg";
        break;
    }
  }

  _updateStyles() {
    this._rotateSprite();
    this._element.style.left = `${this._x}px`;
    this._element.style.top = `${this._y}px`;
    this._element.style.width = `${this._width}px`
    this._element.style.height = `${this._height}px`
  }
}