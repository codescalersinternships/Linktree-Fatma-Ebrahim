const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    historyApiFallback: true,
    proxy: {
      '/': {
        target: 'http://localhost:8000',
        changeOrigin: true,
        ws: false, 
      }
    }
  },
  
})
