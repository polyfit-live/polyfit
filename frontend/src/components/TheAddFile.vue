<template>
  <div>
    <label class="input__label">
      {{label}}
    </label>
    <input
      type="file"
    >
  </div>
</template>

<script >
export default {
  props: {
    label: {
      type: String,
      default: '',
    },
  },
};
</script>

<style scoped>
h1 {
  color: black;
}
</style>
