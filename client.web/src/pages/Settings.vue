<template>
  <div class="content">
    <div class="md-layout">
      <div class="md-layout-item md-medium-size-50 md-xsmall-size-100 md-size-25">
        <h3>{{config.id}}</h3>
        <h3>{{config.datetime_settings}}</h3>
        <h3>{{config.network_settings}}</h3>
        <h3>{{config.audio_settings}}</h3>
        <h3>{{config.radio_settings}}</h3>
        <h3>{{config.alarm_settings}}</h3>
        <h3>{{config.general_settings}}</h3>
      </div>
    </div>
  </div>
</template>

<script>
import { apiClient } from "../main";
export default {
  components: {},
  async created() {
    let id = this.$route.params.id;
    let config = await apiClient.fetchConfigs(id);
    this.config = config;
    console.log(this.config);
  },
  data() {
    return {
      config: {}
    };
  }
};
</script>
