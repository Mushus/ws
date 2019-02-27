const pkg = require('./package.json');

const reactVersion = () => {
  if (pkg.dependencies && pkg.dependencies.react) {
    return { version: pkg.dependencies.react.replace(/[^0-9.]/g, '') };
  }
  if (pkg.devDependencies && pkg.devDependencies.react) {
    return { version: pkg.devDependencies.react.replace(/[^0-9.]/g, '') };
  }
};

module.exports = {
  'extends': [
    'eslint:recommended',
    'plugin:react/recommended',
    'plugin:prettier/recommended'
  ],
  'plugins': [
    '@typescript-eslint',
    'prettier',
    'react'
  ],
  'parser': '@typescript-eslint/parser',
  'parserOptions': {
      'sourceType': 'module',
      'ecmaFeatures': {
          'jsx': true
      },
      'project': './tsconfig.json'
  },
  'rules': {
     /* @typescript-eslint/no-unused-vars */
    'no-unused-vars': 'off'
  },
  'settings': {
    'react': {
      ...reactVersion()
    }
  },
  'globals': {
    'window': true,
    'document': true
  }
};
