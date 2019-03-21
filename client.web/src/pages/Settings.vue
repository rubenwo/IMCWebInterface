<template>
  <div class="content">
    <vue-form-generator :schema="schema" :model="model" :options="formOptions"></vue-form-generator>
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
</template>

<script>
import { apiClient } from "../main";
import VueFormGenerator from "vue-form-generator";

export default {
  components: {
    "vue-form-generator": VueFormGenerator.component
  },
  async created() {
    let id = this.$route.params.id;
    let config = await apiClient.fetchConfig(id);
    this.config = config;
    console.log(this.config);
  },
  data() {
    return {
      config: {},
      model: {
        id: 0,
        name: "John Doe",
        password: "J0hnD03!x4",
        skills: ["Javascript", "VueJS"],
        email: "john.doe@gmail.com",
        factory_reset: false
      },
      schema: {
        fields: [
          {
            type: "input",
            inputType: "text",
            label: "ID (disabled text field)",
            model: "id",
            readonly: true,
            disabled: true
          },
          {
            type: "input",
            inputType: "text",
            label: "Name",
            model: "name",
            placeholder: "Your name",
            featured: true,
            required: true
          },
          {
            type: "input",
            inputType: "password",
            label: "Password",
            model: "password",
            min: 6,
            required: true,
            hint: "Minimum 6 characters",
            validator: VueFormGenerator.validators.string
          },
          {
            type: "select",
            label: "Skills",
            model: "skills",
            values: ["Javascript", "VueJS", "CSS3", "HTML5"]
          },
          {
            type: "input",
            inputType: "email",
            label: "E-mail",
            model: "email",
            placeholder: "User's e-mail address"
          },
          {
            type: "checkbox",
            label: "Factory Reset",
            model: "factory_reset",
            default: false
          }
        ]
      },
      formOptions: {
        validateAfterLoad: true,
        validateAfterChanged: true,
        validateAsync: true
      }
    };
  }
};
</script>
