import Vue from "vue";
import App from "./App.vue";
import Button from "@/components/Button.vue";

import "./css/style.css";
import router from "./router";

Vue.config.productionTip = false;
Vue.component("FrmButton", Button);

new Vue({
  router,

  render: function(h) {
    return h(App);
  }
}).$mount("#app");
