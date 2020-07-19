import Vue from "vue";
import App from "./App.vue";
import Button from "@/components/Button.vue";

import "./css/style.css";

Vue.config.productionTip = false;
Vue.component("FrmButton", Button);

new Vue({
  render: function(h) {
    return h(App);
  },
}).$mount("#app");
