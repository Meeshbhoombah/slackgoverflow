const path        = require('path')
,     webpack     = require('webpack')
,     merge       = require('webpack-merge');


const baseConfig  = require('./webpack.base.config.js');


module.exports = merge(baseConfig, {
  mode: 'development',

  devtool: 'cheap-module-eval-sourcemap',
    
  output: {
    filename    :  'bundle.js',
    path        :  path.resolve(__dirname, 'dist'),
    publicPath  :  '/'
  },

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NamedModulesPlugin()
  ]
});

