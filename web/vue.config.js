'use strict'

const path = require('path')
const buildConf = require('./build.config')
const packageConf = require('./package.json')

function resolve(dir) {
    return path.join(__dirname, dir)
}
module.exports = {
    // Basic configuration details
    publicPath: './',
    outputDir: 'dist',
    assetsDir: 'static',
    lintOnSave: process.env.NODE_ENV === 'development',
    productionSourceMap: false,
    devServer: {
        port: process.env.VUE_APP_CLI_PORT,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        },
        proxy: {
            // Proxy of Key's path to Target location
            // detail: https://cli.vuejs.org/config/#devserver-proxy
            [process.env.VUE_APP_BASE_API]: { //Need a proxy path   E.g '/api'
                target: `${process.env.VUE_APP_BASE_PATH}:${process.env.VUE_APP_SERVER_PORT}/`, //Agent to the target path
                changeOrigin: true,
                pathRewrite: { // Modify path data
                    ['^' + process.env.VUE_APP_BASE_API]: '' // Example '^/api:""' Remove / API string in the path
                }
            }
        },
    },
    configureWebpack: {
        //    @Path walks src folder
        resolve: {
            alias: {
                '@': resolve('src')
            }
        }
    },
    chainWebpack(config) {
        // set preserveWhitespace
        config.module
            .rule('vue')
            .use('vue-loader')
            .loader('vue-loader')
            .tap(options => {
                options.compilerOptions.preserveWhitespace = true
                return options
            })
            .end()
        config
        // https://webpack.js.org/configuration/devtool/#development
            .when(process.env.NODE_ENV === 'development',
            config => config.devtool('cheap-source-map')
        )

        config
            .when(process.env.NODE_ENV !== 'development',
                config => {

                    // Not packaged begin
                    // 1.Currently tested [vue,axios,echarts]Can cdn reference， Other component tests can continue to add
                    // 2.After adding not packaged, you need to public/index.html head Add corresponding CDN resource links
                    config.set('externals', buildConf.cdns.reduce((p, a) => {
                        p[a.name] = a.scope 
                        return p
                    },{}))
                    // Not packaged end

                    config.plugin('html')
                        .tap(args => {
                            if(buildConf.title) {
                                args[0].title = buildConf.title
                            }
                            if(buildConf.cdns.length > 0) {
                                args[0].cdns = buildConf.cdns.map(conf => {
                                    if (conf.path) {
                                        conf.js = `${buildConf.baseCdnUrl}${conf.path}`
                                    } else {
                                        conf.js = `${buildConf.baseCdnUrl}/${conf.name}/${packageConf.dependencies[conf.name].replace('^', '')}/${conf.name}.min.js`
                                    }

                                    return conf
                                })
                            }
                            return args
                        })

                    config
                        .plugin('ScriptExtHtmlWebpackPlugin')
                        .after('html')
                        .use('script-ext-html-webpack-plugin', [{
                            // `runtime` must same as runtimeChunk name. default is `runtime`
                            inline: /single\..*\.js$/
                        }])
                        .end()
                    config
                        .optimization.splitChunks({
                            chunks: 'all',
                            cacheGroups: {
                                libs: {
                                    name: 'chunk-libs',
                                    test: /[\\/]node_modules[\\/]/,
                                    priority: 10,
                                    chunks: 'initial' // only package third parties that are initially dependent
                                },
                                elementUI: {
                                    name: 'chunk-elementUI', // split elementUI into a single package
                                    priority: 20, // the weight needs to be larger than libs and app or it will be packaged into libs or app
                                    test: /[\\/]node_modules[\\/]_?element-ui(.*)/ // in order to adapt to cnpm
                                },
                                commons: {
                                    name: 'chunk-commons',
                                    test: resolve('src/components'), // can customize your rules
                                    minChunks: 3, //  minimum common number
                                    priority: 5,
                                    reuseExistingChunk: true
                                }
                            }
                        })
                    config.optimization.runtimeChunk('single')
                }
            )
    }
}