<script>
    import { onMount } from "svelte";
  import "../styles/app.css";

  let currentAxiom = null
  let isLoading = true

  const fetchAxiom = async() => {
    try {
      const response = await fetch('http://localhost:8080/api/axioms/roulette');
      currentAxiom = await response.json();
    } finally {
      isLoading = false
    }
  }

  onMount(async () => {
   await fetchAxiom();
  })
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