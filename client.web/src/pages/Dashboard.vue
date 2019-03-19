<template>
  <div class="content">
    <div class="md-layout" v-for="config in configs">
      <div class="md-layout-item md-medium-size-50 md-xsmall-size-100 md-size-25">
        <a v-bind:href="'#/settings/'+config.id">
          <stats-card data-background-color="green">
            <template slot="header">
              <md-icon>store</md-icon>
            </template>

            <template slot="content">
              <p class="category">Revenue</p>
              <h3 class="title">$34,245</h3>
            </template>

            <template slot="footer">
              <div class="stats">
                <h4>Device: {{config.id}}</h4>
              </div>
            </template>
          </stats-card>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { StatsCard } from "@/components";
import { apiClient } from "../main";
export default {
  components: {
    StatsCard
  },
  async created() {
  //  let credentials = await apiClient.login("Ruben", "test");
    let configs = await apiClient.fetchConfigs(123);
    this.configs.push(configs);
    console.log(this.configs);
  },
  data() {
    return {
      configs: []
    };
  }
};
</script>
