module.exports = {
  publicPath: "/",
  lintOnSave: process.env.NODE_ENV !== "production",
  configureWebpack: config => {
    if (process.env.NODE_ENV === "production") {
      config.externals = {
        // global app config object
        config: JSON.stringify({ apiUrl: "https://secund.us" })
      };
    } else {
      config.externals = {
        // global app config object
        config: JSON.stringify({ apiUrl: "http://localhost:3000" })
      };
    }
  },
  pwa: {
    name: "Secundus"
  }
};
