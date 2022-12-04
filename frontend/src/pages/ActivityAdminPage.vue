<template>
  <div>
    <h1>Новое упражнение</h1>
    <form class="login__form" action="/profile" @submit="isAutorized">
      <TheInput type="text" label="Название упражнения:" @update:model-value="(newValue)=>(userEmail=newValue)" />
      <TheInput type="text" label="Уровень сложности:" @update:model-value="(newValue)=>(userPassword=newValue)" />
      <TheInput type="text" label="Количество повторений:" @update:model-value="(newValue)=>(userEmail=newValue)" />
      <TheInput type="text" label="Затраченные калории:" @update:model-value="(newValue)=>(userPassword=newValue)" />
      <TheSelect label="Оборудование:" />
      <TheSelect label="Добавить обложку:" />
      <TheAddFile label="Пол:" />
      <TheCheckbox type="checkbox" label="Показывать таймер" @update:model-value="(newValue)=>(userSave=newValue)" />
      <!-- <button type="submit">
            Сохранить
      </button> -->
      <router-link to="/profile" class="login__button">
        Сохранить
      </router-link>
    </form>
  </div>
</template>

<script>
import TheInput from '@/components/TheInput.vue';
import TheCheckbox from '@/components/TheCheckbox.vue';
import TheSelect from '@/components/TheSelect.vue';
import TheAddFile from '@/components/TheAddFile.vue';

export default {
  components: {
    TheInput,
    TheCheckbox,
    TheSelect,
    TheAddFile,
  },
};
</script>

<style scoped>
  h1 {
    padding-top: 60px;
  }
</style>
