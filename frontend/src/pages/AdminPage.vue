<template>
  <section class="admin">
    <div class="admin__wrapper wrapper animate__animated animate__fadeIn">
      <AdminNavigate />
      <div class="admin__home">
        Главная страница
      </div>
    </div>
  </section>
</template>



<script >
import AdminNavigate from '@/components/AdminNavigate.vue';

export default {
  components: {
    AdminNavigate,
  },
  data() {
    return {
      pageActive: 'home',
    };
  },
  methods: {
  },
};

</script>

<style>

.admin{
  min-height: 100vh;
  background-color: #ebedf0;
}

.admin__wrapper{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  max-width: 1440px;
  margin: 0 auto;
  padding: 24px 40px 60px;
  padding-top: 120px;
}

</style>
