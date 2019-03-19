import DashboardLayout from "@/pages/Layout/DashboardLayout.vue";

import Dashboard from "@/pages/Dashboard.vue";
import Settings from "@/pages/Settings.vue";

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: Dashboard
      }, {
        path: "settings/:id",
        name: "Settings",
        component: Settings,
        props: true
      },
    ]
  }
];

export default routes;
