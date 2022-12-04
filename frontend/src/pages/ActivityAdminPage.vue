<template>
  <div>
    <h1>Новое упражнение</h1>
    <form class="login__form" action="/profile" @submit="isAutorized">
      <TheInput type="text" label="Название упражнения:" @update:model-value="(newValue)=>(userEmail=newValue)" />
      <TheInput type="text" label="Уровень сложности:" @update:model-value="(newValue)=>(userPassword=newValue)" />
      <TheInput type="text" label="Количество повторений:" @update:model-value="(newValue)=>(userEmail=newValue)" />
      <TheInput type="text" label="Затраченные калории:" @update:model-value="(newValue)=>(userPassword=newValue)" />
      <TheCheckbox type="checkbox" label="Показывать таймер" @update:model-value="(newValue)=>(userSave=newValue)" />
      <router-link to="/profile" class="login__button">
        Войти
      </router-link>
    </form>
  </div>
</template>

<script>
import TheInput from '@/components/TheInput.vue';
import TheCheckbox from '@/components/TheCheckbox.vue';
export default {
  components: {
    TheInput,
    TheCheckbox,
  },
};
</script>

<style scoped>
  h1 {
    padding-top: 60px;
  }
</style>
