import { Hono } from 'hono'
import { renderer } from './renderer'
import { authApi } from './api/auth/login'

const app = new Hono()

app.use(renderer)

app.get('/', (c) => {
  return c.render(<h1>Hello!</h1>)
})

app.route("/auth", authApi);

export default app
