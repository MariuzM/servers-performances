import Fastify from 'fastify';

const app = Fastify();

const PORT = Number(process.env.PORT) || 3005;

app.get('/', async () => {
  return { hello: 'world' };
});

app.listen({ port: PORT, host: '0.0.0.0' }, (err) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }
  console.log(`Server running on port:${PORT}`);
});
