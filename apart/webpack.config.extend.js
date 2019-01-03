// const path = require('path');

module.exports = config => {
  // uglify の optimize によってハングするので一時的に止める
  config.plugins = config.plugins.filter(
    plugin => plugin.constructor.name !== 'UglifyJsPlugin'
  );
  return config;
};
