const routes = [
    {
        path: '/',
        name: 'in-app-layout',
        children: [
            {
                path: '',
                name: 'home',
                component: () => import('../pages/index/IndexPage.vue')
            }
        ],
        component: () => import('../layouts/Layout.vue')
    }
]
export { routes }