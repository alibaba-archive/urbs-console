import { IConfig } from 'umi-types'; // ref: https://umijs.org/config/

const getPublicPath = () => {
  let publicPath = '/';
  let base = '/'
  if (process.env.BUILD_DEST) {
    const BUILD_GIT_GROUP = process.env.BUILD_GIT_GROUP;
    const BUILD_GIT_PROJECT = process.env.BUILD_GIT_PROJECT;
    const buildArgv = require('yargs-parser')(process.env.BUILD_ARGV_STR);
    if (buildArgv['def_publish_env'] === 'prod') {
      publicPath = `https://g.alicdn.com/${BUILD_GIT_GROUP}/${BUILD_GIT_PROJECT}/`;
      base = 'urbs/'
    } else {
      publicPath = `https://dev.g.alicdn.com/${BUILD_GIT_GROUP}/${BUILD_GIT_PROJECT}/`;
    }
  }
  return { publicPath, base };
};

const config: IConfig = {
  treeShaking: true,
  ...getPublicPath(),
  routes: [
    {
      path: '/',
      component: '../layouts/index',
      routes: [
        {
          path: '/',
          component: '../pages/index',
        },
        {
          path: '/group',
          component: '../pages/group/index',
        },
        {
          path: '/group/:id',
          component: '../pages/group/[id]',
        },
        {
          path: '/user',
          component: '../pages/user/index',
        },
        {
          path: '/user/:name',
          component: '../pages/user/[id]',
        },
        {
          path: '/products/:name',
          component: '../pages/products/[id]',
        },
        {
          path: '/products/:name/tag',
          component: '../pages/products/tag/index',
        },
        {
          path: '/products/:name/module',
          component: '../pages/products/module/index',
        },
        {
          path: '/products/:name/setting',
          component: '../pages/products/setting/index',
        },
        {
          path: '/help',
          component: '../pages/help/index',
        },
      ],
    },
  ],
  plugins: [
    // ref: https://umijs.org/plugin/umi-plugin-react.html
    [
      'umi-plugin-react',
      {
        antd: true,
        dva: true,
        dynamicImport: false,
        title: 'web',
        dll: false,
        locale: true,
        routes: {
          exclude: [
            /models\//,
            /services\//,
            /model\.(t|j)sx?$/,
            /service\.(t|j)sx?$/,
            /components\//,
          ],
        },
      },
    ],
  ],
};
export default config;
