<template>
  <div class="columns is-mobile is-centered mt-6">
    <div class="column is-half">
      <h1 class="title is-1">Login</h1>
      <form @submit.prevent="handleSubmit">
        <div class="field">
          <label class="label">Username</label>
          <div class="control has-icons-left has-icons-right">
            <input
              type="text"
              v-model="username"
              name="username"
              class="input"
              :class="{ 'is-danger': submitted && !username }"
            />
            <span class="icon is-small is-left">
              <i class="fas fa-user"></i>
            </span>
            <span class="icon is-small is-right">
              <i class="fas fa-check"></i>
            </span>
          </div>
          <p class="help is-danger" v-show="submitted && !username">
            Username is required
          </p>
        </div>

        <div class="field">
          <label class="label">Password</label>
          <div class="control has-icons-left has-icons-right">
            <input
              type="password"
              v-model="password"
              name="password"
              class="input"
              :class="{ 'is-danger': submitted && !password }"
            />
            <span class="icon is-small is-left">
              <i class="fas fa-user"></i>
            </span>
            <span class="icon is-small is-right">
              <i class="fas fa-check"></i>
            </span>
          </div>
          <p class="help is-danger" v-show="submitted && !password">
            Password is required
          </p>
        </div>

        <div class="field is-grouped">
          <div class="control">
            <button class="button is-link" :disabled="loggingIn">Submit</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
      submitted: false
    };
  },
  computed: {
    loggingIn() {
      return this.$store.state.authentication.status.loggingIn;
    }
  },
  created() {
    //debugger // eslint-disable-line
    // reset login status
    this.$store.dispatch("authentication/logout");
  },
  mounted() {
    // set page title
    document.title = "Secundus Login";
  },
  methods: {
    handleSubmit() {
      this.submitted = true;
      const { username, password } = this;
      const { dispatch } = this.$store;
      if (username && password) {
        dispatch("authentication/login", { username, password });
      }
    }
  }
};
</script>
