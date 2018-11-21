import Vue from 'vue';
import { VueRoute } from 'vue-router'
import { NuxtAxiosInstance } from 'axios';

declare module '*.vue' {
  const _default: Vue
  export default _default
}

declare module 'vue/types/vue' {
  interface Vue {
    $axios: AxiosInstance,
    $route: VueRoute,
  }
}
