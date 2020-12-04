<template>
  <div>
    <b-sidebar
      id="sidebar"
      type="is-light"
      :fullheight="false"
      :fullwidth="false"
      :overlay="false"
      :right="true"
      v-model="sidebarOpen"
    >
      <div class="p-1">
        <b-menu>
          <b-menu-list label="Menu">
            <b-menu-item
              icon="information-outline"
              tag="router-link"
              type="is-link"
              to="/pitch"
              label="Pitch"
            ></b-menu-item>
            <b-menu-item icon="settings">
              <template slot="label" slot-scope="props">
                Administrator
                <b-icon
                  class="is-pulled-right"
                  :icon="props.expanded ? 'menu-down' : 'menu-up'"
                ></b-icon>
              </template>
              <b-menu-item
                icon="account"
                tag="router-link"
                type="is-link"
                to="/users"
                label="Users"
              ></b-menu-item>
            </b-menu-item>
          </b-menu-list>
          <b-menu-list label="Actions">
            <b-menu-item
              icon="cog-outline"
              tag="router-link"
              type="is-link"
              to="/settings"
              label="Settings"
            ></b-menu-item>
            <b-menu-item
              icon="logout"
              tag="router-link"
              type="is-link"
              to="/login"
              label="Logout"
            ></b-menu-item>
          </b-menu-list>
        </b-menu>
      </div>
    </b-sidebar>
    <b-button
      id="sidebar-toggle-button"
      @click="sidebarOpen = true"
      icon="page-layout-sidebar-right"
    ></b-button>
    <div>
      <div :class="`mx-3 notification is-${alert.type}`" v-if="alert.message">
        <button class="delete"></button>
        {{ alert.message }}
      </div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "app",
  computed: {
    alert() {
      return this.$store.state.alert;
    }
  },
  data() {
    return {
      sidebarOpen: false
    };
  },
  mounted() {
    // set page title
    document.title = "Secundus";

    // set 'app-background' class to body tag
    let bodyElement = document.body;
    bodyElement.classList.add("app-background");

    // check for active theme
    let htmlElement = document.documentElement;
    let theme = localStorage.getItem("theme");

    if (theme === "dark") {
      htmlElement.setAttribute("theme", "dark");
      this.darkMode = true;
    } else {
      htmlElement.setAttribute("theme", "light");
      this.darkMode = false;
    }
  },
  watch: {
    $route() {
      // clear alert on location change
      this.$store.dispatch("alert/clear");
    }
  }
};
</script>

<style lang="scss">
.p-1 {
  padding: 1em;
}
#sidebar {
  .sidebar-content {
    top: 50px;
  }
}
#sidebar-toggle-button {
  position: fixed;
  right: 0px;
  left: auto;
  top: 0px;
  margin: 5px;
}
</style>
