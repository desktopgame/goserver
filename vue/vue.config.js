const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  outputDir: "../web/dist",
  pages: {
    main: {
      entry: "src/MainPage.ts",
      template: "public/MainPage.html",
      filename: "../views/MainPage.html",
    },
  },
  transpileDependencies: true,
});