tailwind.config = {
  theme: {
    extend: {
      colors: {
        brown: '#8f6140',
      }
    }
  }
}

const spawnHorses = () => {
  const horseProps = {
    width: 60,
    height: 60,
    sprite: '/assets/static/images/running_horse.gif',
    screenWidth: window.innerWidth,
    screenHeight: window.innerHeight
  }

  const horses = [
    new Horse({...horseProps, x: 2, y: window.innerHeight - horseProps.height}),
    new Horse({...horseProps, x: window.innerWidth - horseProps.width, direction: HORSE_DIRECTIONS.LEFT}),
    // new Horse({...horseProps, y: (window.innerHeight / 2) - horseProps.height, direction: HORSE_DIRECTIONS.DOWN}),
    // new Horse({...horseProps, x: window.innerWidth - horseProps.height, y: (window.innerHeight / 2) + horseProps.height, direction: HORSE_DIRECTIONS.UP}),
  ]

  for(const horse of horses) {
    document.body.appendChild(horse.htmlElement())
  }

  startAnimationLoop(horses);
}

const startAnimationLoop = (horses) => {
  setInterval(() => {
    for(const horse of horses) {
      horse.update();
    }
  }, 35);
}

const autoplayAudio = () => {
  const element = document.getElementById("audio-autoplay-element");
  element.play();
  document.removeEventListener('click', autoplayAudio);
}

function initialize() {
  spawnHorses();
  document.addEventListener('click', autoplayAudio);
}

