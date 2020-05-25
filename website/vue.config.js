// vue.config.js
module.exports = {
  outputDir: '../static', // 构建输出目录
  publicPath: "/leveldb_admin/static/",
  assetsDir: 'assets',
  productionSourceMap: false,
  devServer: {
    proxy: {
      '': {
        target: 'http://127.0.0.1:4333',
      },
    }
  }
};
