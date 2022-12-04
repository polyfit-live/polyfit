<template>
  <div>
    <label>{{label}}</label>
    <select v-model="selected">
      <option v-for="option in options" :key="option.text" :value="option.value">
        {{option.text}}
      </option>
    </select>
  </div>
</template>

<script>
export default {
  props: {
    label: {
      type: String,
      default: '',
    },
  },

  data: function() {
    return {
      selected: 'A',
      options: [
        { text: 'One', value: 'One' },
        { text: 'Two', value: 'Two' },
        { text: 'Three', value: 'Three' },
      ],
    };
  },
};
</script>

<style  scoped>
h1 {
  color: black;
}
</style>
