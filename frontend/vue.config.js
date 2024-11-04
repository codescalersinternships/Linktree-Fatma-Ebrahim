const { defineConfig } = require('@vue/cli-service');
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    historyApiFallback: true,
    proxy: {
      '/linktree': { 
        target: 'http://localhost:8000', 
        changeOrigin: true,
        ws: false,
      }
    }
  },
});
