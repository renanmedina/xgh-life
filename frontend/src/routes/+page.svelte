<script>
  import { onMount } from "svelte";
  import "../styles/app.css";
  import Horse, {HORSE_DIRECTIONS} from "$lib/horse";

  let currentAxiom = null
  let isLoading = true
  const horses = [];

  const fetchAxiom = async() => {
    try {
      const response = await fetch('/api/axioms/roulette');
      currentAxiom = await response.json();
    } finally {
      isLoading = false
    }
  }

  const spawnHorses = () => {
    const horseProps = {
      width: 60,
      height: 60,
      sprite: '/images/running_horse.gif',
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight
    }

    horses.push(
      new Horse({...horseProps, x: 2, y: window.innerHeight - horseProps.height}),
      new Horse({...horseProps, x: window.innerWidth - horseProps.width, direction: HORSE_DIRECTIONS.LEFT}),
      // new Horse({...horseProps, y: (window.innerHeight / 2) - horseProps.height, direction: HORSE_DIRECTIONS.DOWN}),
      // new Horse({...horseProps, x: window.innerWidth - horseProps.height, y: (window.innerHeight / 2) + horseProps.height, direction: HORSE_DIRECTIONS.UP}),
    )

    for(const horse of horses) {
      document.body.appendChild(horse.htmlElement())
    }
  }

  onMount(async () => {
   await fetchAxiom();
   spawnHorses();
  })

  setInterval(() => {
    for(const horse of horses) {
      horse.update();
    }
  }, 35)
</script>

<style>
  .horse-bubble {
    position: relative;
    max-width: 520px;
  }

  .horse-bubble::before {
    position: absolute;
    content: '';
    left: -30px;
    top: calc(45% - 5px);
    width: 15px;
    height: 15px;
    border-left: 15px solid transparent;
    border-right: 15px solid rgb(253 186 116 / var(--tw-bg-opacity));;
    border-top: 15px solid transparent;
    border-bottom: 15px solid transparent;
  }
</style>

<section class="flex justify-center place-items-center h-full">
  <div class="flex justify-center flex-column">
    <img src="images/xgh_icon.png" alt="xgh horse icon" class="w-32 rounded-full block place-self-center">
    <div class="mr-4 ml-6 rounded-lg bg-orange-300 p-4 block max-w-96 horse-bubble text-yellow-950">
      {#if currentAxiom}
        <h3 class="text-lg font-bold mb-4">{currentAxiom.number} - { currentAxiom.title }</h3>
        { currentAxiom.description }
      {/if}
    </div>
  </div>
</section>