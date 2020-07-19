<template>
  <button class="btn"><slot /></button>
</template>

<style scoped>
.btn {
  background-color: white;
  opacity: 0.9;
  color: #a4508b;
  border: none;
  padding: 10px 25px;
  cursor: pointer;
  font-size: 18px;
  font-weight: bold;
  letter-spacing: 2px;
  text-transform: uppercase;
  border-radius: 1000px;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);

  transition-duration: 100ms;
  transition-property: box-shadow, transform;

  transform: none;
}
.btn:hover {
  /* transform: translateY(-2px); */
  box-shadow: 0px 5px 15px rgba(0, 0, 0, 0.4);
}
</style>
