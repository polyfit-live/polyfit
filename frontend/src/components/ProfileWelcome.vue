<template>
  <div class="profile__welcome">
    <p class="profile__welcome-text">
      {{(12>hourse>=6)?
        ['π', 'π', 'π', 'β', 'π€'][Math.floor(Math.random() * ['π', 'π', 'π', 'β', 'π€'].length)]+' ΠΠΎΠ±ΡΠΎΠ΅ ΡΡΡΠΎ, ':
        (12<=hourse<18)?
          ['π', 'π', 'π', 'π'][Math.floor(Math.random() * ['π', 'π', 'π', 'π'].length)]+' ΠΠΎΠ±ΡΡΠΉ Π΄Π΅Π½Ρ, ':
          (18<=hourse<24)?['π', 'π', 'π', 'π' ][Math.floor(Math.random() * ['π', 'π', 'π', 'π' ].length)]+' ΠΠΎΠ±ΡΡΠΉ Π²Π΅ΡΠ΅Ρ, ':
          ['π', 'π', 'π ', 'π', 'π'][Math.floor(Math.random() * ['π', 'π', 'π ', 'π', 'π'].length)]+' ΠΠΎΠ±ΡΠΎΠΉ Π½ΠΎΡΠΈ, '}}{{name}}!
    </p>
  </div>
</template>

<script>
export default {
  name: 'ProfileWelcomeComponent',
  props: ['name'],
  date() {
    return {
      hours: new Date().getHours(),
    };
  },
};
</script>

<style scoped>

.profile__welcome {
  width: 100%;
  margin-bottom: 36px;
}

.profile__welcome-text {
  font-weight: 500;
  font-size: 32px;
  line-height: 1.1;
}
</style>
