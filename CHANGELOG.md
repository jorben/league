# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.5] - 2024-07-29
### :sparkles: New Features
- [`6d8135e`](https://github.com/jorben/league/commit/6d8135e097b6b709f1a9f6251e8bb57de238ca4e) - 完成用户组页面功能 *(commit by [@jorben](https://github.com/jorben))*
- [`a137ad7`](https://github.com/jorben/league/commit/a137ad7c55c6432a66c1f7eeca93ef6a5bcce4ab) - 完成接口管理，支持新增、编辑、删除接口 *(commit by [@jorben](https://github.com/jorben))*
- [`1033844`](https://github.com/jorben/league/commit/1033844665baf3d8a27e53f7b854502f3b9b1ff8) - 完成管理后台接口权限页面及功能 *(commit by [@jorben](https://github.com/jorben))*
- [`c189673`](https://github.com/jorben/league/commit/c1896730542769a5ec49b43afcd381f32e80ee3a) - 完成管理后台菜单设置页面及菜单项增、删、改功能 *(commit by [@jorben](https://github.com/jorben))*

### :bug: Bug Fixes
- [`d2b3578`](https://github.com/jorben/league/commit/d2b357828cb8e0dfb80e6831fd47fdeba63829a2) - 修复新增接口必填数据未填写提交报错问题 *(commit by [@jorben](https://github.com/jorben))*

### :wrench: Chores
- [`a0ab726`](https://github.com/jorben/league/commit/a0ab726fb588d62bd01130cd11558dd109c6d7bd) - 增加用户组页面，加入用户组支持指定是否新增用户组 *(commit by [@jorben](https://github.com/jorben))*
- [`affd4d5`](https://github.com/jorben/league/commit/affd4d5221fdce184f18fa95039f155038f4da30) - 增加用户组页面，加入用户组支持指定是否新增用户组 *(commit by [@jorben](https://github.com/jorben))*
- [`73473e9`](https://github.com/jorben/league/commit/73473e9a7a31401343c25b6c62c8b8ea84a10342) - 部分完成用户组页面数据展示 *(commit by [@jorben](https://github.com/jorben))*
- [`03bdc4d`](https://github.com/jorben/league/commit/03bdc4def92b50c73c22250a218024fc1dd597d7) - 挑战接口页文件命名，新增权限管理相关页面 *(commit by [@jorben](https://github.com/jorben))*
- [`42e5177`](https://github.com/jorben/league/commit/42e51775cd18fa0fb1233964472db41ae023d69d) - 增加获取权限规则接口，接口权限页面展示各分组规则列表 *(commit by [@jorben](https://github.com/jorben))*
- [`f398e26`](https://github.com/jorben/league/commit/f398e26778dd1fb52a2280a442180d602a7db3d1) - 管理后台菜单设置页面完成数据展示 *(commit by [@jorben](https://github.com/jorben))*
- [`d6c20c8`](https://github.com/jorben/league/commit/d6c20c8f659b434ebfdd10efe30c9f036264927e) - 管理后台菜单设置页完成数据编辑、删除功能 *(commit by [@jorben](https://github.com/jorben))*


## [0.0.4] - 2024-07-16
### :sparkles: New Features
- [`5597de8`](https://github.com/jorben/league/commit/5597de85b4758bb3d1d7b62ba81d978f5f75024f) - 增加用户基本信息接口，管理后台头部调用用户信息 *(commit by [@jorben](https://github.com/jorben))*
- [`1d22894`](https://github.com/jorben/league/commit/1d22894d8ee17d0125fe572ee63482497b27452d) - 串通登录、回调流程 *(commit by [@jorben](https://github.com/jorben))*
- [`dc07405`](https://github.com/jorben/league/commit/dc07405fce72a657ab52b706877e958377ddd499) - 完成用户列表页数据获取和渲染 *(commit by [@jorben](https://github.com/jorben))*
- [`6a85c33`](https://github.com/jorben/league/commit/6a85c33ff64a837e7ec760133b2a258ed2fef95d) - 增加微信PC扫码登录 *(commit by [@jorben](https://github.com/jorben))*
- [`cac9d3c`](https://github.com/jorben/league/commit/cac9d3c8cf7ee28c0887f233c69124b8c24f2e53) - 完成用户详情页面数据对接，完成用户状态变更功能 *(commit by [@jorben](https://github.com/jorben))*
- [`973d788`](https://github.com/jorben/league/commit/973d788e54f5c99ed707535b0c5618a4843bc128) - 增加解绑登录渠道功能 *(commit by [@jorben](https://github.com/jorben))*
- [`585211c`](https://github.com/jorben/league/commit/585211cfdff3ff0ef1ebe570f95e6208ad9d7248) - 增加登录后回跳来源地址 *(commit by [@jorben](https://github.com/jorben))*
- [`96bf79d`](https://github.com/jorben/league/commit/96bf79db6c2d3551239d8333971555004526942f) - 增加删除用户接口，完成用户删除操作流程对接 *(commit by [@jorben](https://github.com/jorben))*
- [`647c42a`](https://github.com/jorben/league/commit/647c42a147c882fc6dfbe998dd23e1e351d0e6bb) - 增加加入/退出用户组接口，用户管理完成对接用户组操作 *(commit by [@jorben](https://github.com/jorben))*

### :bug: Bug Fixes
- [`d25e2eb`](https://github.com/jorben/league/commit/d25e2eb0485ef6c12bba04a8b09e6e77c66b4439) - 修正login页多次发起请求问题，修正auth/login返回的message内容 *(commit by [@jorben](https://github.com/jorben))*

### :wrench: Chores
- [`5c15e5e`](https://github.com/jorben/league/commit/5c15e5e958285c1efef410c72abf4d12cd837c7b) - 规划logout接口 *(commit by [@jorben](https://github.com/jorben))*
- [`e6f664a`](https://github.com/jorben/league/commit/e6f664aa270a9670f07a1dcdcaeed1fbfcf4cdef) - 搭建前端项目 *(commit by [@jorben](https://github.com/jorben))*
- [`ae5254f`](https://github.com/jorben/league/commit/ae5254f734de631182e4118e5970a1748d1c05c7) - 调整PageNot'Found路由策略 *(commit by [@jorben](https://github.com/jorben))*
- [`6f2aed0`](https://github.com/jorben/league/commit/6f2aed01f72d2150fca9bd964698e317a69892f7) - 搭建好了管理后台页面框架、登录页框架 *(commit by [@jorben](https://github.com/jorben))*
- [`25c13e2`](https://github.com/jorben/league/commit/25c13e291a3e7e15a7a782d9198eb447b21f917e) - 增加前端工程服务及路由 *(commit by [@jorben](https://github.com/jorben))*
- [`2c6a8e3`](https://github.com/jorben/league/commit/2c6a8e3d84e8e072aca3f24c94dc18c4775ad599) - 调整public下静态资源路径 *(commit by [@jorben](https://github.com/jorben))*
- [`bebc4d7`](https://github.com/jorben/league/commit/bebc4d77f87200edbca0facef1f0ddbcccb39133) - 拆分管理后台菜单模块，构建了用户列表页和详情内容框架 *(commit by [@jorben](https://github.com/jorben))*
- [`c3050a7`](https://github.com/jorben/league/commit/c3050a75a1223a1de5bc95c3054379dbe9cf29ef) - 完成用户列表、用户详情页面框架，管理后台菜单从接口获取，后端构建包含前端资源 *(commit by [@jorben](https://github.com/jorben))*
- [`6cc32ae`](https://github.com/jorben/league/commit/6cc32aeb8337e25c0b87ba6727c669a43f6a60d1) - 修增typo，CONSTANTS *(commit by [@jorben](https://github.com/jorben))*
- [`bf84af9`](https://github.com/jorben/league/commit/bf84af91d48d48df4eea7aacaa8a07ac64f61011) - 前端资源开启gzip，登录页随机背景图 *(commit by [@jorben](https://github.com/jorben))*


## [0.0.3] - 2024-07-03
### :sparkles: New Features
- [`8bdcdfa`](https://github.com/jorben/league/commit/8bdcdfa04041dd06acf68c846be5549114e9347e) - Github登录增加state校验 *(commit by [@jorben](https://github.com/jorben))*
- [`be33edf`](https://github.com/jorben/league/commit/be33edfd4a4b7fa13fabc4ad42ef158482ef8669) - 登录接口增加返回用户id和token时间信息 *(commit by [@jorben](https://github.com/jorben))*
- [`b94b8b8`](https://github.com/jorben/league/commit/b94b8b8e6a391f026538b3de36370182c3866e94) - 增加auth/renew接口，支持刷新jwt *(commit by [@jorben](https://github.com/jorben))*

### :bug: Bug Fixes
- [`c6de514`](https://github.com/jorben/league/commit/c6de5141eff2adfd5137ffa7e38c52804b93185f) - changelog config branch not match *(commit by [@jorben](https://github.com/jorben))*
- [`cd5eba8`](https://github.com/jorben/league/commit/cd5eba8957f181540ff6413e30aed15370ce8f2d) - changelog格式错乱，删除重新生成 *(commit by [@jorben](https://github.com/jorben))*

[0.0.3]: https://github.com/jorben/league/compare/0.0.2...0.0.3
[0.0.4]: https://github.com/jorben/league/compare/0.0.3...0.0.4
[0.0.5]: https://github.com/jorben/league/compare/0.0.4...0.0.5
