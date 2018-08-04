const express = require('express')
const request = require('request')

async function main () {
  const app = express()
  const port = process.env.PORT || 3000
  const serviceUrl = process.env.SERVICE_URL

  app.get('/', (req, res) => {
    res.status(200).json({
      message: 'hello world'
    })
  })

  // Liveness probe (health check)
  app.get('/health', (req, res) => {
    res.status(200).json({
      status: 'healthy'
    })
  })

  // livenessProbe: // can be reused for readinessProbe too
  //  httpGet:
  //    path: /health
  //    port: 3000
  // initialDelaySeconds: 30
  // timeoutSeconds: 1

  app.get('/service', (req, res) => {
    let url = req.query.url || serviceUrl
    if (!url) {
      return res.status(400).json({
        status: 'BadRequest',
        message: 'The url provided is invalid'
      })
    }

    request(url, { json: true }, (error, response, body) => {
      if (error) {
        return res.status(400).json({
          url,
          status: 'BadRequest',
          message: error.message
        })
      }
      return res.status(200).json({
        data: {
          body
        }
      })
    })
  })

  // app.listen() returns http.Server
  let server = app.listen(port, () => {
    console.log(`listening to port*:${port}. press ctrl + c to cancel.`)
  })

  return server
}

main()
  .then(server => {
    // Graceful server stop
    process.on('SIGTERM', async () => {
      try {
        await server.close()
        process.exit(0)
      } catch (error) {
        console.error(error)
        process.exit(-1)
      }
    })
  })
  .catch(console.error)
