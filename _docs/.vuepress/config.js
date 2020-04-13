module.exports = {
    base: '/goribot/',
    locales: {
        '/': {
            lang: 'zh-CN',
            title: 'Goribot 使用文档',
            description: '🕷 一个轻量的分布式友好的 Golang 爬虫框架。',
        },
        // '/en/': {
        //     lang: 'en-US',
        //     title: 'Goribot Document',
        //     description: '🕷 A lightweight distributed friendly Golang crawler framework.',
        // },
    },
    head: [
        ['script', { src: "https://www.googletagmanager.com/gtag/js?id=UA-131918267-5", async: true }],
        ['script', {},
            " window.dataLayer = window.dataLayer || [];\
                function gtag(){dataLayer.push(arguments);}\
                gtag('js', new Date());\
                gtag('config', 'UA-131918267-5');"],
        ['link', { rel: 'icon', href: '/favicon.ico' }],
    ],
    // plugins: [
    //     ['@vuepress/google-analytics', { 'ga': 'UA-131918267-5' }]
    // ],
    themeConfig: {
        repo: 'zhshch2002/goribot',
        editLinks: true,
        docsDir: '_docs',
        sidebar: [
            '/',
            '/get-start.md',
            '/distributed.md',
            '/extensions.md',
            '/component.md',
        ],
        sidebarDepth: 2,
        displayAllHeaders: true,
        locales: {
            '/': {
                lang: 'zh-CN',
                selectText: 'Languages',
                label: '简体中文',
                editLinkText: '在 GitHub 上编辑此页',
                lastUpdated: '上次更新于',
            }
        }
    },
}