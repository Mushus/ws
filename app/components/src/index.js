import { component } from './util.js';
import Accordion from './accordion/component.js';
import Sidenav from './sidenav/component.js';
import Simple from './simple.html';

component('x-sidenav', Sidenav);
component('x-accordion', Accordion);
component('x-simple', Simple);
