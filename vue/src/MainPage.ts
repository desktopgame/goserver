import { createApp } from 'vue';
import MainPage from './components/MainPage.vue';
import store from './store';

createApp(MainPage).use(store).mount('#app');
