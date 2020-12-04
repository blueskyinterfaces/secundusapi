<template>
  <div class="container m-6">
    <!-- theme switcher -->
    <input
      type="checkbox"
      id="theme-switch"
      class="theme-switch"
      v-model="darkMode"
    />
    <label for="theme-switch">
      <span v-if="darkMode === true"> Switch to Light Mode </span>
      <span v-else> Switch to Dark Mode </span>
    </label>
  </div>
</template>

<script>
export default {
  name: "Admin",
  data() {
    return {
      darkMode: false,
    };
  },
  mounted() {
    // set page title
    document.title = "Secundus Admin";

    // check for active theme
    let theme = localStorage.getItem("theme");
    if (theme === "dark") {
      this.darkMode = true;
    } else {
      this.darkMode = false;
    }
  },
  watch: {
    darkMode: function () {
      // add/remove class to/from html tag
      let htmlElement = document.documentElement;

      if (this.darkMode) {
        localStorage.setItem("theme", "dark");
        htmlElement.setAttribute("theme", "dark");
      } else {
        localStorage.setItem("theme", "light");
        htmlElement.setAttribute("theme", "light");
      }
    },
  },
};
</script>
