const express = require('express')
const httpProxy = require('express-http-proxy')
const app = express()

const auth_service_proxy = httpProxy('http://service.auth')
const config_service_proxy = httpProxy('http://service.config')
const web_client_proxy = httpProxy('http://client.web')

app.use((req, res, next) => {
  // TODO: my authentication logic
  next()
})

app.post('/auth', (req, res, next) => {
  auth_service_proxy(req, res, next)
})

app.post('/config', (req, res, next) => {
  config_service_proxy(req, res, next)
})

app.all('/config/:id', (req, res, next) => {
  config_service_proxy(req, res, next)
})

app.get('/*', (req, res, next) => {
  web_client_proxy(req, res, next)
})

app.listen(80, () => {
  console.log("Started api-gateway on port 80...")
})