
<template>
  <div class="animate__animated animate__fadeIn">
    <section class="workouts-settings">
      <WorkoutsSettings @change-sort="ChangeSort" @change-type="ChangeType" @change-range="ChangeTime" @change-difficulty="ChangeDifficulty" @change-search="ChangeSearch" />
    </section>
    <section class="workouts-content">
      <div class="workouts-content__wrapper">
        <ul ref="parent" class="workouts-content__list">
          <li v-for="(workout, index) in apiWorkouts" :key="index" class="workouts-content__item">
            <WorkoutsItem :workout="workout" />
          </li>
        </ul>
      </div>
    </section>
  </div>
</template>

<script>
import WorkoutsItem from '@/components/WorkoutsItem';
import WorkoutsSettings from '@/components/WorkoutsSettings';
import axios from 'axios';

export default {
  components: {
    WorkoutsItem,
    WorkoutsSettings,
  },
  data() {
    return {
      id: 0,
      varSort: 'По новизне',
      varType: 'all',
      varTime: [0, 67],
      varSearch: '',
      varDifficulty: 'all',
      apiWorkouts: [],
      workouts: [{
        image: require('@/assets/WorkoutImage-0.png'),
        title: 'Работа над резервной копией',
        type: 'Силовая',
        calories: 113,
        difficulty: 'средне',
        time: 22,
        rec: true,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Спортивный коврик.',
      },
      {
        id: 1,
        image: require('@/assets/WorkoutImage-1.png'),
        title: 'Тренируем нашу выносливость',
        type: 'Скоростная',
        calories: 140,
        difficulty: 'легко',
        time: 25,
        rec: true,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',

      },
      {
        id: 2,
        image: require('@/assets/WorkoutImage-2.png'),
        title: 'Тренировка на различные группы мышц',
        type: 'Скоростная',
        calories: 260,
        difficulty: 'сложно',
        time: 30,
        rec: true,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Спортивный коврик.',

      },
      {
        id: 3,
        image: require('@/assets/WorkoutImage-3.png'),
        title: 'Бубновый валет',
        type: 'Выносливость',
        calories: 132,
        difficulty: 'легко',
        time: 23,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 4,
        image: require('@/assets/WorkoutImage-4.png'),
        title: 'Мощь! И не только',
        type: 'Выносливость',
        calories: 288,
        difficulty: 'сложно',
        time: 33,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 5,
        image: require('@/assets/WorkoutImage-5.png'),
        title: 'Безграничные возможности',
        type: 'Силовая',
        calories: 180,
        difficulty: 'средне',
        time: 29,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Спортивный коврик.',
      },
      {
        id: 6,
        image: require('@/assets/WorkoutImage-6.png'),
        title: 'На благо своего здоровья',
        type: 'Скоростная',
        calories: 102,
        difficulty: 'средне',
        time: 16,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 7,
        image: require('@/assets/WorkoutImage-7.png'),
        title: 'Стремление к усовершенствованию',
        type: 'Силовая',
        calories: 174,
        difficulty: 'сложно',
        time: 35,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 8,
        image: require('@/assets/WorkoutImage-8.png'),
        title: 'Улучшаем свои характеристики',
        type: 'Скоростная',
        calories: 94,
        difficulty: 'легко',
        time: 20,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 9,
        image: require('@/assets/WorkoutImage-9.png'),
        title: 'Равномерная тренировка с направленностью на ноги',
        type: 'Выносливость',
        calories: 110,
        difficulty: 'средне',
        time: 18,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 10,
        image: require('@/assets/WorkoutImage-10.png'),
        title: 'Утро бывает добрым',
        type: 'Силовая',
        calories: 112,
        difficulty: 'легко',
        time: 21,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 11,
        image: require('@/assets/WorkoutImage-11.png'),
        title: 'На все тело',
        type: 'Силовая',
        calories: 380,
        difficulty: 'средне',
        time: 67,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 12,
        image: require('@/assets/WorkoutImage-12.png'),
        title: 'Утепляемся сами!',
        type: 'Выносливость',
        calories: 263,
        difficulty: 'сложно',
        time: 46,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 13,
        image: require('@/assets/WorkoutImage-13.png'),
        title: 'Держим свое тело в тонусе',
        type: 'Силовая',
        calories: 140,
        difficulty: 'легко',
        time: 24,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 14,
        image: require('@/assets/WorkoutImage-14.png'),
        title: 'Тренировка, которая вгоняет В ритм жизни',
        type: 'Скоростная',
        calories: 116,
        difficulty: 'средне',
        time: 23,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      },
      {
        id: 15,
        image: require('@/assets/WorkoutImage-15.png'),
        title: 'Динамика и энергия',
        type: 'Выносливость',
        calories: 292,
        difficulty: 'средне',
        time: 46,
        rec: false,
        description: 'Данная тренировка расчитана на улучшение выносливости тела к умеренным физическим нагрузкам. Время, выделенное на выполнение упражнений соответсвует уровню сложности данной тренировки.',
        inventory: 'Гантели, спортивный коврик.',
      }],
    };
  },
  computed: {
    sortWorkouts() {
      let variant = this.varSort;
      let type = this.varType;
      let time = this.varTime;
      let text = this.varSearch;
      let difficulty = this.varDifficulty;
      let list = this.sortByType(this.apiWorkouts, type);
      list = this.sortByTime(list, time);
      list = this.sortByDifficulty(list, difficulty);
      list = this.sortBySearch(list, text);
      if (variant === 'По новизне') { return list; }
      return list.slice().sort(function(a, b) {
        { return ((variant === 'По длительности') ? (a.time > b.time) : (a.cal > b.cal)) ? 1 : -1; }
      });
    },
  },
  created() {
    let url = '/api/workout';
    axios({
      method: 'get', // you can set what request you want to be
      url: url,
    })
      .then(response => (this.apiWorkouts = response.data));
  },
  methods: {
    sortByType(list, type) {
      if (type === 'power') { return list.filter(function(item) { return item.type === 'Силовая'; }); }
      if (type === 'stamina') { return list.filter(function(item) { return item.type === 'Выносливость'; }); }
      if (type === 'speed') { return list.filter(function(item) { return item.type === 'Скоростная'; }); }
      return list;
    },
    sortByDifficulty(list, difficulty) {
      if (difficulty === 'Легко') { return list.filter(function(item) { return item.difficulty === 'легко'; }); }
      if (difficulty === 'Средне') { return list.filter(function(item) { return item.difficulty === 'средне'; }); }
      if (difficulty === 'Сложно') { return list.filter(function(item) { return item.difficulty === 'сложно'; }); }
      return list;
    },
    sortByTime(list, time) {
      return list.filter(function(item) { return (item.time <= time[1]) & (item.time >= time[0]); });
    },
    sortBySearch(list, text) {
      // return list.filter(function(item) { return item.difficulty === 'сложно'; });
      return list.filter(function(item) {
        return (item.name.toLowerCase().includes(text.toLowerCase()) & item.desc.toLowerCase().includes(text.toLowerCase()));
      });
    },
    ChangeSort(variant) {
      this.varSort = variant;
    },
    ChangeType(variant) {
      this.varType = variant;
    },
    ChangeDifficulty(difficulty) {
      this.varDifficulty = difficulty;
    },
    ChangeTime(minTime, maxTime) {
      this.varTime = [minTime, maxTime];
    },
    ChangeSearch(search) {
      this.varSearch = search;
    },
  },
};
</script>

<style scoped>
.workouts-content__wrapper {
  max-width: 1440px;
  margin: 0 auto;
  padding: 24px 40px 60px;
}

.workouts-content__list {
  position: relative;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
  list-style: none;
}
</style>
