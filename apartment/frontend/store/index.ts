import Vue from 'vue'
import Vuex from 'vuex'
import app from './modules/app'

export default function store() {
  return new Vuex.Store({
    modules: {
      app,
    },
  })
}
