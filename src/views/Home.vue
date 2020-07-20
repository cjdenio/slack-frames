<template>
  <div class="home page">
    <div class="centered">
      <div
        :style="{
          border: '5px solid rgba(0, 0, 0, 0.2)',
          borderRadius: '10px',
          boxShadow: '0px 5px 15px rgba(0, 0, 0, 0.2)',
          textAlign: 'center',
          padding: '30px',
        }"
      >
        <img class="logo" src="../assets/logo.png" />
        <transition name="main" mode="out-in" appear>
          <Loader v-if="!loaded" />
          <div v-if="loaded && authed">
            <h2>Welcome, {{ name }}!</h2>
            <FrmButton to="/about" icon="arrow-right">Let's go</FrmButton>
          </div>
          <a href="/api/login" v-if="authed === false"
            ><img
              alt="Sign in with Slack"
              height="40"
              width="172"
              src="https://platform.slack-edge.com/img/sign_in_with_slack.png"
              srcset="
                https://platform.slack-edge.com/img/sign_in_with_slack.png    1x,
                https://platform.slack-edge.com/img/sign_in_with_slack@2x.png 2x
              "
          /></a>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Loader from "@/components/Loader.vue";
import axios from "axios";

export default {
  name: "Home",
  components: {
    Loader,
  },
  data: () => ({
    loaded: false,
    authed: null,
    name: "",
  }),
  async created() {
    try {
      const resp = await axios.get("/api/user");
      this.loaded = true;
      this.authed = true;
      this.name = resp.data.real_name_normalized;
    } catch (e) {
      this.loaded = true;
      this.authed = false;
    }
  },
};
</script>

<style>
.centered {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.logo {
  width: 450px;
  display: block;
}

.main-enter-active,
.main-leave-active {
  transition: opacity 0.5s;
}
.main-enter, .main-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
