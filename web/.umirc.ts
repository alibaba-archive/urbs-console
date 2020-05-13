import { IConfig } from 'umi-types'; // ref: https://umijs.org/config/
const defEnv = process.env.NODE_DEF_ENV;
console.log('defEnv:', defEnv);
const config: IConfig = {
  treeShaking: true,
  hash: true,
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
