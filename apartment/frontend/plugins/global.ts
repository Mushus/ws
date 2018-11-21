import Vue from 'vue'
import Breadcrumb from '~/components/shared/Breadcrumb.vue';
import Container from '~/components/shared/Container.vue';
import Content from '~/components/shared/Content.vue';
import InputItem from '~/components/shared/InputItem.vue';
import Subheader from '~/components/shared/Subheader.vue';

import DateInput from '~/components/inputs/DateInput.vue';
import InsertButton from '~/components/inputs/InsertButton.vue';

//shared
Vue.component('Breadcrumb', Breadcrumb);
Vue.component('Container', Container);
Vue.component('Content', Content);
Vue.component('InputItem', InputItem);
Vue.component('Subheader', Subheader);

// inputs
Vue.component('DateInput', DateInput);
Vue.component('InsertButton', InsertButton);
